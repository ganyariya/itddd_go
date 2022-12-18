package interactor

import (
	"fmt"

	"github.com/ganyariya/itddd_go/pkg/boundary/input"
	"github.com/ganyariya/itddd_go/pkg/boundary/irepository"
	"github.com/ganyariya/itddd_go/pkg/boundary/output"
	"github.com/ganyariya/itddd_go/pkg/domain/entity"
	"github.com/ganyariya/itddd_go/pkg/domain/service"
	"github.com/ganyariya/itddd_go/pkg/domain/value"
)

type UserCreateInteractor struct {
	userRepository irepository.IUserRepository
	userService    *service.UserService
}

func NewUserCreateInteractor(userRepo irepository.IUserRepository, userServiece *service.UserService) *UserCreateInteractor {
	return &UserCreateInteractor{
		userRepository: userRepo,
		userService:    userServiece,
	}
}

func (u *UserCreateInteractor) Handle(inputData *input.UserCreateInputData) (*output.UserCreateOutputData, error) {
	userId, err := value.NewUserId()
	if err != nil {
		return nil, err
	}
	user := entity.NewUser(userId, inputData.UserName)

	if u.userService.Exists(user) {
		return nil, fmt.Errorf("already user exists")
	}

	err = u.userRepository.Save(user)
	if err != nil {
		return nil, err
	}
	return &output.UserCreateOutputData{User: user}, nil
}
