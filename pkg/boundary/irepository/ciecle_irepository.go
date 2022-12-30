package irepository

import (
	"github.com/ganyariya/itddd_go/pkg/domain/entity"
	"github.com/ganyariya/itddd_go/pkg/domain/value"
)

type ICircleRepository interface {
	Save(*entity.Circle)
	FindByCircleId(value.CircleId) (*entity.Circle, error)
	FindByCircleName(value.CircleName) (*entity.Circle, error)
}
