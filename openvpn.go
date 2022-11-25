// Package openvpn provides an API for connecting to, receiving messages from, and sending commands
// to an OpenVPN management socket.
package openvpn

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/cedws/openvpn-mgmt-go/message"
)

const (
	messageClientConnect     = "CLIENT:CONNECT"
	messageClientReauth      = "CLIENT:REAUTH"
	messageClientEnv         = "CLIENT:ENV"
	messageClientEstablished = "CLIENT:ESTABLISHED"
	messageClientDisconnect  = "CLIENT:DISCONNECT"
	messageEcho              = "ECHO"
	messageFatal             = "FATAL"
	messageHold              = "HOLD"
	messageInfo              = "INFO"
	messageLog               = "LOG"
)

type (
	OnClientConnectFunc     func(message.ClientConnect)
	OnClientReauthFunc      func(message.ClientReauth)
	OnClientEstablishedFunc func(message.ClientEstablished)
	OnClientDisconnectFunc  func(message.ClientDisconnect)
	OnEchoFunc              func(message.Echo)
	OnFatalFunc             func(message.Fatal)
	OnHoldFunc              func(message.Hold)
	OnInfoFunc              func(message.Info)
	OnLogFunc               func(message.Log)
)

type Command interface {
	String() string
}

type Event interface {
	Parse(string) error
}

type Socket struct {
	onClientConnect     OnClientConnectFunc
	onClientReauth      OnClientReauthFunc
	onClientEstablished OnClientEstablishedFunc
	onClientDisconnect  OnClientDisconnectFunc
	onEcho              OnEchoFunc
	onFatal             OnFatalFunc
	onHold              OnHoldFunc
	onInfo              OnInfoFunc
	onLog               OnLogFunc
	conn                io.ReadWriteCloser
	messageCh           chan any
}

func DialUnix(addr string) (*Socket, error) {
	sock, err := net.Dial("unix", addr)
	if err != nil {
		return nil, err
	}

	return &Socket{
		conn:      sock,
		messageCh: make(chan any),
	}, nil
}

func DialTCP(addr string) (*Socket, error) {
	sock, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Socket{
		conn:      sock,
		messageCh: make(chan any),
	}, nil
}

func parse[T any, U interface {
	*T
	Event
}](line string) U {
	u := U(new(T))
	_ = u.Parse(line)
	return u
}

func readEnv(messageCh <-chan any) map[string]string {
	envMap := make(map[string]string)

	for m := range messageCh {
		if envVar, ok := m.(*message.ClientEnvVar); ok {
			if envVar.End {
				break
			}

			envMap[envVar.Key] = envVar.Value
		} else {
			panic("expected to read environment")
		}
	}

	return envMap
}

func (s *Socket) OnClientConnect(handleFunc OnClientConnectFunc) {
	s.onClientConnect = handleFunc
}

func (s *Socket) OnClientReauth(handleFunc OnClientReauthFunc) {
	s.onClientReauth = handleFunc
}

func (s *Socket) OnClientEstablished(handleFunc OnClientEstablishedFunc) {
	s.onClientEstablished = handleFunc
}

func (s *Socket) OnClientDisconnect(handleFunc OnClientDisconnectFunc) {
	s.onClientDisconnect = handleFunc
}

func (s *Socket) OnEcho(handleFunc OnEchoFunc) {
	s.onEcho = handleFunc
}

func (s *Socket) OnFatal(handleFunc OnFatalFunc) {
	s.onFatal = handleFunc
}

func (s *Socket) OnHold(handleFunc OnHoldFunc) {
	s.onHold = handleFunc
}

func (s *Socket) OnInfo(handleFunc OnInfoFunc) {
	s.onInfo = handleFunc
}

func (s *Socket) OnLog(handleFunc OnLogFunc) {
	s.onLog = handleFunc
}

func (s *Socket) Start() {
	go s.read()

	for m := range s.messageCh {
		if m == nil {
			continue
		}

		switch m := m.(type) {
		case *message.ClientConnect:
			m.Env = readEnv(s.messageCh)
			if s.onClientConnect != nil {
				go s.onClientConnect(*m)
			}
		case *message.ClientDisconnect:
			m.Env = readEnv(s.messageCh)
			if s.onClientDisconnect != nil {
				go s.onClientDisconnect(*m)
			}
		case *message.ClientEstablished:
			m.Env = readEnv(s.messageCh)
			if s.onClientEstablished != nil {
				go s.onClientEstablished(*m)
			}
		case *message.ClientReauth:
			m.Env = readEnv(s.messageCh)
			if s.onClientReauth != nil {
				go s.onClientReauth(*m)
			}
		case *message.Echo:
			if s.onEcho != nil {
				go s.onEcho(*m)
			}
		case *message.Fatal:
			if s.onFatal != nil {
				go s.onFatal(*m)
			}
		case *message.Hold:
			if s.onHold != nil {
				go s.onHold(*m)
			}
		case *message.Info:
			if s.onInfo != nil {
				go s.onInfo(*m)
			}
		case *message.Log:
			if s.onLog != nil {
				go s.onLog(*m)
			}
		case message.ClientEnvVar:
			panic("got unexpected client env var message")
		}
	}
}

func (s *Socket) read() {
	scanner := bufio.NewScanner(s.conn)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 || line[0] != '>' {
			continue
		}

		n := line[1:]
		args := strings.SplitN(n, ",", 2)
		if len(args) < 1 {
			continue
		}

		switch source := args[0]; source {
		case messageClientConnect:
			s.messageCh <- parse[message.ClientConnect](line)
		case messageClientDisconnect:
			s.messageCh <- parse[message.ClientDisconnect](line)
		case messageClientEstablished:
			s.messageCh <- parse[message.ClientEstablished](line)
		case messageClientReauth:
			s.messageCh <- parse[message.ClientReauth](line)
		case messageEcho:
			s.messageCh <- parse[message.Echo](line)
		case messageFatal:
			s.messageCh <- parse[message.Fatal](line)
		case messageHold:
			s.messageCh <- parse[message.Hold](line)
		case messageInfo:
			s.messageCh <- parse[message.Info](line)
		case messageLog:
			s.messageCh <- parse[message.Log](line)
		case messageClientEnv:
			s.messageCh <- parse[message.ClientEnvVar](line)
		}
	}
}

func (s *Socket) Close() error {
	close(s.messageCh)
	return s.conn.Close()
}

func (s *Socket) Dispatch(c Command) error {
	if _, err := s.conn.Write([]byte(c.String() + "\n")); err != nil {
		return fmt.Errorf("openvpn: error dispatching command: %w", err)
	}

	return nil
}
