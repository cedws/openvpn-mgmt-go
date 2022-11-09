package command

import (
	"fmt"
	"strconv"
)

type NeedStr struct {
	Type string
	Str  string
}

func (n NeedStr) String() string {
	return fmt.Sprintf("needstr %v %v", strconv.Quote(n.Type), strconv.Quote(n.Str))
}
