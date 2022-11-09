package command

type Help struct{}

func (h Help) String() string {
	return "help"
}
