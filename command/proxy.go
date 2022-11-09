package command

import (
	"fmt"
)

type Proxy struct {
	Type             string
	Host             string
	Port             int
	CleartextAllowed bool
}

func (p Proxy) String() string {
	nct := "nct"
	if p.CleartextAllowed {
		nct = ""
	}
	if p.Host != "" {
		return fmt.Sprintf("proxy %v %v %v %v", p.Type, p.Host, p.Port, nct)
	}
	return fmt.Sprintf("proxy %v", p.Type)
}
