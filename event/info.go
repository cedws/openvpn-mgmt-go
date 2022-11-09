package event

import (
	"fmt"
	"strings"
)

type Info struct {
	Message string
}

func (i *Info) Parse(line string) (err error) {
	if !strings.HasPrefix(line, ">INFO") {
		return fmt.Errorf("invalid message")
	}

	args := strings.SplitN(line, ":", 2)
	if len(args) < 2 {
		return fmt.Errorf("malformed message, not enough args")
	}

	*i = Info{args[1]}
	return
}
