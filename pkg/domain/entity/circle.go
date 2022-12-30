package entity

import "github.com/ganyariya/itddd_go/pkg/domain/value"

type Circle struct {
	Id      value.CircleId
	Name    value.CircleName
	Owner   value.UserId
	Members []value.UserId
}

func NewCircle(Id value.CircleId, Name value.CircleName, Owner value.UserId, Members []value.UserId) *Circle {
	return &Circle{
		Id:      Id,
		Name:    Name,
		Owner:   Owner,
		Members: Members,
	}
}
