package command

import (
	"fmt"
)

type Hold struct {
	Flag string
}

func (h Hold) String() string {
	return fmt.Sprintf("hold %v", h.Flag)
}
