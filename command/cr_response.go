package command

import (
	"fmt"
)

type CrResponse struct {
	Response string
}

func (c CrResponse) String() string {
	return fmt.Sprintf("cr-response %v", c.Response)
}
