package stub

import (
	"fmt"

	"github.com/ganyariya/itddd_go/pkg/boundary/irepository"
	"github.com/ganyariya/itddd_go/pkg/domain/entity"
	"github.com/ganyariya/itddd_go/pkg/domain/value"
)

type CircleRepository struct {
	circles []*entity.Circle
}

func NewCircleRepository() irepository.ICircleRepository {
	return &CircleRepository{circles: []*entity.Circle{}}
}

func (cr *CircleRepository) Save(circle *entity.Circle) {
	cr.circles = append(cr.circles, circle)
}

func (cr *CircleRepository) FindByCircleId(circleId value.CircleId) (*entity.Circle, error) {
	for _, c := range cr.circles {
		if c.Id == circleId {
			return c, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func (cr *CircleRepository) FindByCircleName(circleName value.CircleName) (*entity.Circle, error) {
	for _, c := range cr.circles {
		if c.Name == circleName {
			return c, nil
		}
	}
	return nil, fmt.Errorf("not found")
}
