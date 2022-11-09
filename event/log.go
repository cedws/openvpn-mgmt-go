package event

import (
	"fmt"
	"strconv"
	"strings"
)

type Log struct {
	Timestamp int64
	Flags     string
	Message   string
}

func (l *Log) Parse(line string) (err error) {
	if !strings.HasPrefix(line, ">LOG") {
		return fmt.Errorf("invalid message")
	}

	args := strings.SplitN(line, ":", 2)
	if len(args) < 2 {
		return fmt.Errorf("malformed message, not enough args")
	}
	args = strings.SplitN(args[1], ",", 3)
	if len(args) < 3 {
		return fmt.Errorf("malformed message, not enough args")
	}

	time, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return fmt.Errorf("invalid timestamp: %w", err)
	}

	*l = Log{time, args[1], args[2]}
	return
}
