package entity

import "github.com/ganyariya/itddd_go/pkg/domain/value"

type User struct {
	Id        value.UserId
	Name      string
	IsPremium bool
}

func NewUser(id value.UserId, name string, isPremium bool) *User {
	return &User{Id: id, Name: name, IsPremium: isPremium}
}

func (u *User) Equals(o *User) bool {
	return u.Id == o.Id
}

func (u *User) ChangeName(name string) {
	u.Name = name
}
