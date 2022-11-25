package message

import (
	"fmt"
	"strings"
)

type Fatal struct {
	Message string
}

func (f *Fatal) Parse(line string) (err error) {
	if !strings.HasPrefix(line, ">FATAL") {
		return ErrMalformed
	}

	args := strings.SplitN(line, ":", 2)
	if len(args) < 2 {
		return fmt.Errorf("expected at least 2 args")
	}

	*f = Fatal{args[1]}
	return
}
