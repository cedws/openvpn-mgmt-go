package message

import (
	"fmt"
)

var (
	ErrMalformed     = fmt.Errorf("message is malformed")
	ErrMissingFields = fmt.Errorf("message is missing fields")
)
