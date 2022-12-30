package dto

import "github.com/ganyariya/itddd_go/pkg/domain/entity"

type CircleDTO struct {
	Id        string
	Name      string
	OwnerId   string
	MemberIds []string
}

func NewCircleDTOByCircle(circle *entity.Circle) *CircleDTO {
	memberIds := []string{}
	for _, id := range circle.MemberIds {
		memberIds = append(memberIds, id.Id)
	}
	return &CircleDTO{
		Id:        circle.Id.Id,
		Name:      circle.Name.Name,
		OwnerId:   circle.OwnerId.Id,
		MemberIds: memberIds,
	}
}
