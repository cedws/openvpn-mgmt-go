package command

import (
	"fmt"
)

type Verb struct {
	Mode int
}

func (v Verb) String() string {
	if v.Mode != 0 {
		return fmt.Sprintf("verb %v", v.Mode)
	}
	return "verb"
}
