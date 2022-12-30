package entity

import "github.com/ganyariya/itddd_go/pkg/domain/value"

type Circle struct {
	Id        value.CircleId
	Name      value.CircleName
	OwnerId   value.UserId
	MemberIds []value.UserId
}

func NewCircle(Id value.CircleId, Name value.CircleName, OwnerId value.UserId, MemberIds []value.UserId) *Circle {
	return &Circle{
		Id:        Id,
		Name:      Name,
		OwnerId:   OwnerId,
		MemberIds: MemberIds,
	}
}
