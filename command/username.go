package command

import (
	"fmt"
	"strconv"
)

type Username struct {
	Type     string
	Username string
}

func (u Username) String() string {
	return fmt.Sprintf("username %v %v", strconv.Quote(u.Type), strconv.Quote(u.Username))
}
