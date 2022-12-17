package irepository

import (
	"github.com/ganyariya/itddd_go/pkg/domain/entity"
	"github.com/ganyariya/itddd_go/pkg/domain/value"
)

type IUserRepository interface {
	Save(*entity.User) error
	Find(value.UserId) (*entity.User, error)
}
