package command

import (
	"fmt"
)

type ClientAuthNt struct {
	ClientID int64
	KeyID    int64
}

func (c ClientAuthNt) String() string {
	return fmt.Sprintf("client-auth-nt %v %v", c.ClientID, c.KeyID)
}
