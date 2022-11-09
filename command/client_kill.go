package command

import (
	"fmt"
)

type ClientKill struct {
	ClientID int64
}

func (c ClientKill) String() string {
	return fmt.Sprintf("client-kill %v", c.ClientID)
}
