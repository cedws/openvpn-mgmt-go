package command

import (
	"fmt"
	"strconv"
)

type ClientPendingAuth struct {
	ClientID int64
	Extra    string
	Timeout  int
}

func (c ClientPendingAuth) String() string {
	return fmt.Sprintf("client-pending-auth %v %v %v", c.ClientID, strconv.Quote(c.Extra), c.Timeout)
}
