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

type CircleJoinCommand struct {
	UserId   string
	CircleId string
}

func NewCircleJoinCommand(UserId string, CircleId string) *CircleJoinCommand {
	return &CircleJoinCommand{
		UserId:   UserId,
		CircleId: CircleId,
	}
}
