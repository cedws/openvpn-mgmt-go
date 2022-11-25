package message

import (
	"fmt"
	"strings"
)

type Info struct {
	Message string
}

func (i *Info) Parse(line string) (err error) {
	if !strings.HasPrefix(line, ">INFO") {
		return ErrMalformed
	}

	args := strings.SplitN(line, ":", 2)
	if len(args) < 2 {
		return fmt.Errorf("expected at least 2 args")
	}

	*i = Info{args[1]}
	return
}
