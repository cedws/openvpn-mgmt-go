// Package openvpn provides an API for connecting to, receiving messages from, and sending commands
// to an OpenVPN management socket.
package openvpn

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/cedws/openvpn-mgmt-go/event"
)

const (
	MessageClientConnect     = "CLIENT:CONNECT"
	MessageClientReauth      = "CLIENT:REAUTH"
	MessageClientEnv         = "CLIENT:ENV"
	MessageClientEstablished = "CLIENT:ESTABLISHED"
	MessageClientDisconnect  = "CLIENT:DISCONNECT"
	MessageEcho              = "ECHO"
	MessageFatal             = "FATAL"
	MessageHold              = "HOLD"
	MessageInfo              = "INFO"
	MessageLog               = "LOG"
)

type Command interface {
	String() string
}

type Socket struct {
	conn       io.ReadWriteCloser
	handleFunc HandleFunc
	errorFunc  ErrorFunc
	messageCh  chan any
}

type HandleFunc func(any)
type ErrorFunc func(error)

func DialUnix(addr string) (*Socket, error) {
	sock, err := net.Dial("unix", addr)
	if err != nil {
		return nil, err
	}

	return &Socket{sock, nil, nil, make(chan any)}, nil
}

func DialTCP(addr string) (*Socket, error) {
	sock, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Socket{sock, nil, nil, make(chan any)}, nil
}

func (s *Socket) HandleFunc(handleFunc HandleFunc) {
	s.handleFunc = handleFunc
}

func (s *Socket) ErrorFunc(errorFunc ErrorFunc) {
	s.errorFunc = errorFunc
}

func (s *Socket) onEvent(e any) {
	if s.handleFunc != nil {
		go s.handleFunc(e)
	}
}

func (s *Socket) onError(e error) {
	if s.errorFunc != nil {
		go s.errorFunc(e)
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

		// bleh, is there really no better way to do this?
		switch source := args[0]; source {
		case MessageClientConnect:
			v := &event.ClientConnect{}
			if err := v.Parse(line); err != nil {
				s.onError(fmt.Errorf("openvpn: error handling %v: %w", MessageClientConnect, err))
			}
			s.messageCh <- v
		case MessageClientReauth:
			v := &event.ClientReauth{}
			if err := v.Parse(line); err != nil {
				s.onError(fmt.Errorf("openvpn: error handling %v: %w", MessageClientReauth, err))
			}
			s.messageCh <- v
		case MessageClientEnv:
			v := &event.ClientEnvVar{}
			if err := v.Parse(line); err != nil {
				s.onError(fmt.Errorf("openvpn: error handling %v: %w", MessageClientEnv, err))
			}
			s.messageCh <- v
		case MessageClientEstablished:
			v := &event.ClientEstablished{}
			if err := v.Parse(line); err != nil {
				s.onError(fmt.Errorf("openvpn: error handling %v: %w", MessageClientEstablished, err))
			}
			s.messageCh <- v
		case MessageClientDisconnect:
			v := &event.ClientDisconnect{}
			if err := v.Parse(line); err != nil {
				s.onError(fmt.Errorf("openvpn: error handling %v: %w", MessageClientDisconnect, err))
			}
			s.messageCh <- v
		case MessageEcho:
			v := &event.Echo{}
			if err := v.Parse(line); err != nil {
				s.onError(fmt.Errorf("openvpn: error handling %v: %w", MessageEcho, err))
			}
			s.messageCh <- v
		case MessageFatal:
			v := &event.Fatal{}
			if err := v.Parse(line); err != nil {
				s.onError(fmt.Errorf("openvpn: error handling %v: %w", MessageFatal, err))
			}
			s.messageCh <- v
		case MessageHold:
			v := &event.Hold{}
			if err := v.Parse(line); err != nil {
				s.onError(fmt.Errorf("openvpn: error handling %v: %w", MessageHold, err))
			}
			s.messageCh <- v
		case MessageInfo:
			v := &event.Info{}
			if err := v.Parse(line); err != nil {
				s.onError(fmt.Errorf("openvpn: error handling %v: %w", MessageInfo, err))
			}
			s.messageCh <- v
		case MessageLog:
			v := &event.Log{}
			if err := v.Parse(line); err != nil {
				s.onError(fmt.Errorf("openvpn: error handling %v: %w", MessageLog, err))
			}
			s.messageCh <- v
		}
	}
}

func (s *Socket) readEnv() map[string]string {
	envMap := make(map[string]string)

	for m := range s.messageCh {
		if envVar, ok := m.(*event.ClientEnvVar); ok {
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

func (s *Socket) Start() {
	go s.read()

	for message := range s.messageCh {
		switch m := message.(type) {
		case *event.ClientConnect:
			m.Env = s.readEnv()
		case *event.ClientDisconnect:
			m.Env = s.readEnv()
		case *event.ClientEstablished:
			m.Env = s.readEnv()
		case *event.ClientReauth:
			m.Env = s.readEnv()
		}

		s.onEvent(message)
	}
}

func (s *Socket) Close() error {
	return s.conn.Close()
}

func (s *Socket) Dispatch(c Command) error {
	if _, err := s.conn.Write([]byte(c.String() + "\n")); err != nil {
		return fmt.Errorf("openvpn: error dispatching command: %w", err)
	}

	return nil
}
