package command

import (
	"fmt"
)

type Signal struct {
	Type string
}

func (s Signal) String() string {
	return fmt.Sprintf("signal %v", s.Type)
}
