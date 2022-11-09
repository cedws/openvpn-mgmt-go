package command

type Exit struct{}

func (e Exit) String() string {
	return "exit"
}
