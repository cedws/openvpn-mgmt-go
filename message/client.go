package message

import (
	"fmt"
	"strconv"
	"strings"
)

type ClientConnect struct {
	ClientID int64
	KeyID    int64
	Env      map[string]string
}

func (c *ClientConnect) Parse(line string) error {
	if !strings.HasPrefix(line, ">CLIENT:CONNECT") {
		return ErrMalformed
	}

	args := strings.SplitN(line, ",", 3)
	if len(args) < 3 {
		return fmt.Errorf("expected at least 3 args")
	}

	cid, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return fmt.Errorf("invalid client id: %w", err)
	}

	kid, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		return fmt.Errorf("invalid key id: %w", err)
	}

	*c = ClientConnect{cid, kid, nil}
	return nil
}

type ClientReauth struct {
	ClientID int64
	KeyID    int64
	Env      map[string]string
}

func (c *ClientReauth) Parse(line string) error {
	if !strings.HasPrefix(line, ">CLIENT:REAUTH") {
		return ErrMalformed
	}

	args := strings.SplitN(line, ",", 3)
	if len(args) < 3 {
		return fmt.Errorf("expected at least 3 args")
	}

	cid, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return fmt.Errorf("invalid client id: %w", err)
	}

	kid, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		return fmt.Errorf("invalid key id: %w", err)
	}

	*c = ClientReauth{cid, kid, nil}
	return nil
}

type ClientEnvVar struct {
	End   bool
	Key   string
	Value string
}

func (c *ClientEnvVar) Parse(line string) error {
	if !strings.HasPrefix(line, ">CLIENT:ENV") {
		return ErrMalformed
	}

	if line == ">CLIENT:ENV,END" {
		*c = ClientEnvVar{true, "", ""}
		return nil
	}

	args := strings.SplitN(line, ",", 2)
	if len(args) < 2 {
		return fmt.Errorf("expected at least 2 args")
	}

	kv := strings.SplitN(args[1], "=", 2)
	if len(kv) < 2 {
		return fmt.Errorf("malformed key value pair")
	}

	*c = ClientEnvVar{false, kv[0], kv[1]}
	return nil
}

type ClientEstablished struct {
	ClientID int64
	Env      map[string]string
}

func (c *ClientEstablished) Parse(line string) (err error) {
	if !strings.HasPrefix(line, ">CLIENT:ESTABLISHED") {
		return ErrMalformed
	}

	args := strings.SplitN(line, ",", 2)
	if len(args) < 2 {
		return fmt.Errorf("expected at least 2 args")
	}

	cid, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return fmt.Errorf("invalid client id: %w", err)
	}

	*c = ClientEstablished{cid, nil}
	return
}

type ClientDisconnect struct {
	ClientID int64
	Env      map[string]string
}

func (c *ClientDisconnect) Parse(line string) (err error) {
	if !strings.HasPrefix(line, ">CLIENT:DISCONNECT") {
		return ErrMalformed
	}

	args := strings.SplitN(line, ",", 2)
	if len(args) < 2 {
		return fmt.Errorf("expected at least 2 args")
	}

	cid, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return fmt.Errorf("invalid client id: %w", err)
	}

	*c = ClientDisconnect{cid, nil}
	return
}

type ClientAddress struct {
	ClientID int64
	Address  string
	Primary  bool
}

func (c *ClientAddress) Parse(line string) (err error) {
	if !strings.HasPrefix(line, ">CLIENT:ADDRESS") {
		return ErrMalformed
	}

	args := strings.SplitN(line, ",", 4)
	if len(args) < 4 {
		return fmt.Errorf("expected at least 4 args")
	}

	cid, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return fmt.Errorf("invalid client id: %w", err)
	}

	p, err := strconv.ParseBool(args[3])
	if err != nil {
		return fmt.Errorf("malformed message, invalid bool: %w", err)
	}

	*c = ClientAddress{cid, args[2], p}
	return
}

type ClientChallengeResponse struct {
	ClientID       int64
	KeyID          int64
	ResponseBase64 string
}

func (c *ClientChallengeResponse) Parse(line string) (err error) {
	if !strings.HasPrefix(line, ">CLIENT:CR_RESPONSE") {
		return ErrMalformed
	}

	args := strings.SplitN(line, ",", 4)
	if len(args) < 4 {
		return fmt.Errorf("expected at least 4 args")
	}

	cid, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return fmt.Errorf("invalid client id: %w", err)
	}

	kid, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		return fmt.Errorf("invalid key id: %w", err)
	}

	*c = ClientChallengeResponse{cid, kid, args[3]}
	return
}
