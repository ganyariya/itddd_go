package service

import (
	"fmt"

	"github.com/ganyariya/itddd_go/pkg/boundary/irepository"
	"github.com/ganyariya/itddd_go/pkg/domain/entity"
	"github.com/ganyariya/itddd_go/pkg/domain/service"
	"github.com/ganyariya/itddd_go/pkg/domain/value"
)

type UserApplicationService struct {
	userRepository irepository.IUserRepository
	userService    service.UserService
}

func NewUserApplicationService(userRepo irepository.IUserRepository, userService service.UserService) *UserApplicationService {
	return &UserApplicationService{userRepository: userRepo, userService: userService}
}

func (uas *UserApplicationService) Register(name string) error {
	userId, err := value.NewUserId()
	if err != nil {
		return err
	}

	user := entity.NewUser(userId, name)
	fmt.Println(user)
	return nil
}
