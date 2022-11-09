package command

import (
	"fmt"
)

type RSASig struct {
	Signature string
}

func (r RSASig) String() string {
	return fmt.Sprintf("rsa-sig\n%v\nEND", r.Signature)
}
