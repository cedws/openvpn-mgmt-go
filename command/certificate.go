package command

import (
	"fmt"
)

type Certificate struct {
	Cert string
}

func (c Certificate) String() string {
	return fmt.Sprintf("certificate\n%v\nEND", c.Cert)
}
