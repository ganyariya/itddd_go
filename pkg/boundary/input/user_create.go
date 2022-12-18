package input

type UserCreateInputData struct {
	UserName string
}

func NewUserCreateInputData(userName string) *UserCreateInputData {
	return &UserCreateInputData{UserName: userName}
}
