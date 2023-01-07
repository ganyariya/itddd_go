package service

import (
	"fmt"

	"github.com/ganyariya/itddd_go/pkg/application/command"
	"github.com/ganyariya/itddd_go/pkg/boundary/dto"
	"github.com/ganyariya/itddd_go/pkg/boundary/irepository"
	"github.com/ganyariya/itddd_go/pkg/domain/factory"
	"github.com/ganyariya/itddd_go/pkg/domain/service"
	"github.com/ganyariya/itddd_go/pkg/domain/value"
)

/*
ドメインサービスはドメインを表現するために必要なもの。
それらドメインで表現された世界の問題を解決するものがアプリケーションサービス。
*/
type UserApplicationService struct {
	userRepository irepository.IUserRepository
	userService    *service.UserService
	userFactory    factory.IUserFactory
}

func NewUserApplicationService(userRepo irepository.IUserRepository, userService *service.UserService, userFactory factory.IUserFactory) *UserApplicationService {
	return &UserApplicationService{userRepository: userRepo, userService: userService, userFactory: userFactory}
}

func (uas *UserApplicationService) Register(command *command.UserRegisterCommand) (*dto.UserDTO, error) {
	user, err := uas.userFactory.Create(command.Name, command.IsPremium)
	if err != nil {
		return nil, err
	}
	if uas.userService.Exists(user) {
		return nil, fmt.Errorf("exists")
	}

	uas.userRepository.Save(user)
	return dto.NewUserDTOByUser(user), nil
}

func (uas *UserApplicationService) Get(command *command.UserGetCommand) (*dto.UserDTO, error) {
	userId, err := value.NewUserIdByIdString(command.Id)
	if err != nil {
		return nil, err
	}

	user, err := uas.userRepository.Find(userId)
	if err != nil {
		return nil, err
	}
	return dto.NewUserDTOByUser(user), nil
}

func (uas *UserApplicationService) Update(command *command.UserUpdateCommand) (*dto.UserDTO, error) {
	userId, err := value.NewUserIdByIdString(command.Id)
	if err != nil {
		return nil, err
	}

	user, err := uas.userRepository.Find(userId)
	if err != nil {
		return nil, err
	}

	if command.Name != nil {
		user.Name = *command.Name
	}
	err = uas.userRepository.Save(user)
	return dto.NewUserDTOByUser(user), err
}

func (uas *UserApplicationService) Delete(command *command.UserDeleteCommand) error {
	userId, err := value.NewUserIdByIdString(command.Id)
	if err != nil {
		return err
	}

	user, err := uas.userRepository.Find(userId)
	if err != nil {
		return err
	}

	uas.userRepository.Delete(user)
	return nil
}

func (uas *UserApplicationService) GetAll() ([]*dto.UserDTO, error) {
	userDtos := []*dto.UserDTO{}
	users, err := uas.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		userDtos = append(userDtos, dto.NewUserDTOByUser(user))
	}
	return userDtos, nil
}
