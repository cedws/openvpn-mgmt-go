package command

import (
	"fmt"
)

type PKCS11IdGet struct {
	Id int
}

func (p PKCS11IdGet) String() string {
	return fmt.Sprintf("pkcs11-id-get %v", p.Id)
}
