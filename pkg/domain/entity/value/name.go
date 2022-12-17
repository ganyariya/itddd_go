package value

type FullName struct {
	FirstName string
	LastName  string
}

func NewFullName(first, last string) FullName {
	return FullName{first, last}
}
