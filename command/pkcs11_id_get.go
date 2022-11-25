package command

import (
	"fmt"
)

type PKCS11IdGet struct {
	ID int
}

func (p PKCS11IdGet) String() string {
	return fmt.Sprintf("pkcs11-id-get %v", p.ID)
}
