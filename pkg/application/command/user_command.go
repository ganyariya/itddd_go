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

type UserRegisterCommand struct {
	Name      string
	IsPremium bool
}

func NewUserRegisterCommand(Name string, IsPremium bool) *UserRegisterCommand {
	return &UserRegisterCommand{Name: Name, IsPremium: IsPremium}
}

type UserGetCommand struct {
	Id string
}

func NewUserGetCommand(Id string) *UserGetCommand {
	return &UserGetCommand{Id: Id}
}

type UserDeleteCommand struct {
	Id string
}

func NewUserDeleteCommand(Id string) *UserDeleteCommand {
	return &UserDeleteCommand{Id: Id}
}
