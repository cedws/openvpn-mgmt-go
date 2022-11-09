package command

import (
	"fmt"
)

type Kill struct {
	Instance string
}

func (k Kill) String() string {
	return fmt.Sprintf("kill %v", k.Instance)
}
