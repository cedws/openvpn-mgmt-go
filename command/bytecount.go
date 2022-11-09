package command

import (
	"fmt"
)

type Bytecount struct {
	Seconds int
}

func (b Bytecount) String() string {
	return fmt.Sprintf("bytecount %v", b.Seconds)
}
