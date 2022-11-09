package command

import (
	"fmt"
	"strconv"
)

type NeedOK struct {
	Type         string
	Confirmation string
}

func (n NeedOK) String() string {
	return fmt.Sprintf("needok %v %v", strconv.Quote(n.Type), strconv.Quote(n.Confirmation))
}
