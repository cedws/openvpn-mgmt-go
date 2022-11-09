package command

import (
	"fmt"
)

type Echo struct {
	Mode string
}

func (e Echo) String() string {
	return fmt.Sprintf("echo %v", e.Mode)
}
