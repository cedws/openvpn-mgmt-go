package command

import (
	"fmt"
)

type Status struct {
	Format int
}

func (s Status) String() string {
	if s.Format != 0 {
		return fmt.Sprintf("status %v", s.Format)
	}
	return "status"
}
