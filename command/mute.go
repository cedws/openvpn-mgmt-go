package command

import (
	"fmt"
)

type Mute struct {
	Mode int
}

func (m Mute) String() string {
	if m.Mode != 0 {
		return fmt.Sprintf("mute %v", m.Mode)
	}
	return "mute"
}
