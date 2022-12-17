package interactor

import (
	"github.com/ganyariya/itddd_go/pkg/boundary/input"
	"github.com/ganyariya/itddd_go/pkg/boundary/irepository"
	"github.com/ganyariya/itddd_go/pkg/boundary/output"
	"github.com/ganyariya/itddd_go/pkg/domain/entity"
	"github.com/ganyariya/itddd_go/pkg/domain/value"
)

type UserCreateInteractor struct {
	userRepository irepository.IUserRepository
}

func NewUserCreateInteractor(userRepo irepository.IUserRepository) *UserCreateInteractor {
	return &UserCreateInteractor{userRepository: userRepo}
}

func (u *UserCreateInteractor) Handle(inputData *input.UserCreateInputData) (*output.UserCreateOutputData, error) {
	userId, err := value.NewUserId(inputData.UserId)
	if err != nil {
		return nil, err
	}
	user := entity.NewUser(userId, inputData.UserName)
	err = u.userRepository.Save(user)
	if err != nil {
		return nil, err
	}
	return &output.UserCreateOutputData{User: user}, nil
}
