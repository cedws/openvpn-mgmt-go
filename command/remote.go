package command

import (
	"fmt"
)

type Remote struct {
	Action string
	Host   string
	Port   int
}

func (r Remote) String() string {
	if r.Host != "" {
		return fmt.Sprintf("remote %v %v %v", r.Action, r.Host, r.Port)
	}
	return fmt.Sprintf("remote %v", r.Action)
}
