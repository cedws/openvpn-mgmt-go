package command

import (
	"fmt"
	"strings"
)

type ClientAuth struct {
	ClientID     int64
	KeyID        int64
	ClientConfig []string
}

func (c ClientAuth) String() string {
	block := strings.Join(c.ClientConfig, "\n")
	return fmt.Sprintf("client-auth %v %v\n%v\nEND", c.ClientID, c.KeyID, block)
}
