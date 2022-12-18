package command

/*
UseCase の Input Output に近い
*/
type UserUpdateCommand struct {
	Id   string
	Name *string
}

func NewUserUpdateCommand(Id string, Name *string) *UserUpdateCommand {
	return &UserUpdateCommand{Id: Id, Name: Name}
}

type UserDeleteCommand struct {
	Id string
}

func NewUserDeleteCommand(Id string) *UserDeleteCommand {
	return &UserDeleteCommand{Id: Id}
}
