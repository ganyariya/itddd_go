package entity

import (
	"fmt"

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

func (c *Circle) countUsers() int {
	return 1 + len(c.MemberIds)
}
func (c *Circle) isFull() bool {
	return c.countUsers() >= CircleMaxMemberNum
}
func (c *Circle) JoinUser(userId value.UserId) error {
	if c.isFull() {
		return fmt.Errorf("already full")
	}
	c.MemberIds = append(c.MemberIds, userId)
	return nil
}
