package command

type ForgetPasswords struct{}

func (f ForgetPasswords) String() string {
	return "forget-passwords"
}
