package entity

import (
	"github.com/ganyariya/itddd_go/pkg/domain/value"
)

const CircleMaxMemberNum = 30

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

func (c *Circle) CountUsers() int {
	return 1 + len(c.MemberIds)
}
func (c *Circle) IsFull(UpperBound int) bool {
	return c.CountUsers() >= UpperBound
}

func (c *Circle) JoinUser(userId value.UserId) error {
	c.MemberIds = append(c.MemberIds, userId)
	return nil
}
