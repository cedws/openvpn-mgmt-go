package event

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
		return fmt.Errorf("invalid message")
	}

	args := strings.SplitN(line, ":", 2)
	if len(args) < 2 {
		return fmt.Errorf("malformed message, not enough args")
	}
	args = strings.SplitN(args[1], ",", 2)
	if len(args) < 2 {
		return fmt.Errorf("malformed message, not enough args")
	}

	time, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return fmt.Errorf("invalid timestamp: %w", err)
	}

	*e = Echo{time, args[1]}
	return
}
