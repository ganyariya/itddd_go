package input

type UserCreateInputData struct {
	UserId   string
	UserName string
}

func NewUserCreateInputData(userId string, userName string) *UserCreateInputData {
	return &UserCreateInputData{UserId: userId, UserName: userName}
}
