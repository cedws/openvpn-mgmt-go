package command

import (
	"fmt"
)

type AuthRetry struct {
	Mode string
}

func (a AuthRetry) String() string {
	return fmt.Sprintf("auth-retry %v", a.Mode)
}
