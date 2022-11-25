package message

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
		return ErrMalformed
	}

	args := strings.SplitN(line, ":", 2)
	if len(args) < 2 {
		return fmt.Errorf("expected at least 2 args")
	}
	args = strings.SplitN(args[1], ",", 3)
	if len(args) < 3 {
		return fmt.Errorf("expected at least 3 args")
	}

	time, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return fmt.Errorf("invalid timestamp: %w", err)
	}

	*l = Log{time, args[1], args[2]}
	return
}
