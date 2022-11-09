package event

import (
	"fmt"
	"strconv"
	"strings"
)

type Hold struct {
	Message string
	Wait    int
}

func (h *Hold) Parse(line string) (err error) {
	if !strings.HasPrefix(line, ">HOLD") {
		return fmt.Errorf("invalid message")
	}

	args := strings.SplitN(line, ":", 3)
	if len(args) < 3 {
		return fmt.Errorf("malformed message, not enough args")
	}

	wait, err := strconv.Atoi(args[2])
	if err != nil {
		return fmt.Errorf("invalid wait time: %w", err)
	}

	*h = Hold{args[1], wait}
	return
}
