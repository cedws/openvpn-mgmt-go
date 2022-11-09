package command

type PKCS11IdCount struct{}

func (p PKCS11IdCount) String() string {
	return "pkcs11-id-count"
}
