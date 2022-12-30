package command

type CircleRegisterCommand struct {
	UserId     string
	CircleName string
}

func NewCircleRegisterCommand(UserId string, CircleName string) *CircleRegisterCommand {
	return &CircleRegisterCommand{
		UserId:     UserId,
		CircleName: CircleName,
	}
}
