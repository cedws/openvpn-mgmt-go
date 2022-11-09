package command

import (
	"fmt"
)

type Log struct {
	Mode string
}

func (l Log) String() string {
	return fmt.Sprintf("log %v", l.Mode)
}
