package command

import (
	"fmt"
)

type State struct {
	Mode string
}

func (s State) String() string {
	return fmt.Sprintf("state %v", s.Mode)
}
