package input

type UserCreateInputData struct {
	UserName  string
	IsPremium bool
}

func NewUserCreateInputData(userName string, isPremium bool) *UserCreateInputData {
	return &UserCreateInputData{UserName: userName, IsPremium: isPremium}
}
