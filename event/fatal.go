package event

import (
	"fmt"
	"strings"
)

type Fatal struct {
	Message string
}

func (f *Fatal) Parse(line string) (err error) {
	if !strings.HasPrefix(line, ">FATAL") {
		return fmt.Errorf("invalid message")
	}

	args := strings.SplitN(line, ":", 2)
	if len(args) < 2 {
		return fmt.Errorf("malformed message, not enough args")
	}

	*f = Fatal{args[1]}
	return
}
