package message

import (
	"fmt"
	"strconv"
	"strings"
)

type Echo struct {
	Timestamp int64
	Message   string
}

func (e *Echo) Parse(line string) (err error) {
	if !strings.HasPrefix(line, ">ECHO") {
		return ErrMalformed
	}

	args := strings.SplitN(line, ":", 2)
	if len(args) < 2 {
		return fmt.Errorf("expected at least 2 args")
	}
	args = strings.SplitN(args[1], ",", 2)
	if len(args) < 2 {
		return fmt.Errorf("expected at least 2 args")
	}

	time, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return fmt.Errorf("invalid timestamp: %w", err)
	}

	*e = Echo{time, args[1]}
	return
}
