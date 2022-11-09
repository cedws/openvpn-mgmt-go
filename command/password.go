package command

import (
	"fmt"
	"strconv"
)

type Password struct {
	Type     string
	Password string
}

func (p Password) String() string {
	return fmt.Sprintf("password %v %v", strconv.Quote(p.Type), strconv.Quote(p.Password))
}
