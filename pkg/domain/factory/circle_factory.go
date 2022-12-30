package factory

import (
	"github.com/ganyariya/itddd_go/pkg/domain/entity"
	"github.com/ganyariya/itddd_go/pkg/domain/value"
)

type ICircleFactory interface {
	Create(string, value.UserId) (*entity.Circle, error)
}

type CircleFactory struct{}

func NewCircleFactory() ICircleFactory {
	return &CircleFactory{}
}

func (cf *CircleFactory) Create(circleNameString string, userId value.UserId) (*entity.Circle, error) {
	circleName, err := value.NewCircleName(circleNameString)
	if err != nil {
		return nil, err
	}
	circleId, err := value.NewCircleId()
	if err != nil {
		return nil, err
	}

	circle := entity.NewCircle(circleId, circleName, userId, []value.UserId{})
	return circle, err
}
