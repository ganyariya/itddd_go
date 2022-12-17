package entity

import "github.com/ganyariya/itddd_go/pkg/domain/value"

type User struct {
	Id   value.UserId
	Name string
}

func NewUser(id value.UserId, name string) *User {
	return &User{Id: id, Name: name}
}

func (u *User) Equals(o *User) bool {
	return u.Id == o.Id
}

func (u *User) ChangeName(name string) {
	u.Name = name
}
