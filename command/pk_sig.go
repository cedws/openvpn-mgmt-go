package command

import (
	"fmt"
)

type PkSig struct {
	Signature string
}

func (p PkSig) String() string {
	return fmt.Sprintf("pk-sig\n%v\nEND", p.Signature)
}
