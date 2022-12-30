package service

import (
	"github.com/ganyariya/itddd_go/pkg/boundary/irepository"
	"github.com/ganyariya/itddd_go/pkg/domain/entity"
)

type CircleService struct {
	circleRepository irepository.ICircleRepository
}

func NewCircleService(circleRepository irepository.ICircleRepository) *CircleService {
	return &CircleService{circleRepository: circleRepository}
}

func (cs *CircleService) Exists(circle *entity.Circle) bool {
	d, _ := cs.circleRepository.FindByCircleName(circle.Name)
	return d != nil
}
