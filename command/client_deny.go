package command

import (
	"fmt"
	"strconv"
)

type ClientDeny struct {
	ClientID     int64
	KeyID        int64
	Reason       string
	ClientReason string
}

func (c ClientDeny) String() string {
	return fmt.Sprintf("client-deny %v %v %v %v", c.ClientID, c.KeyID, strconv.Quote(c.Reason), strconv.Quote(c.ClientReason))
}
