package command

import (
	"fmt"
)

type Version struct {
	Version int
}

func (v Version) String() string {
	if v.Version != 0 {
		return fmt.Sprintf("version %v", v.Version)
	}
	return "version"
}
