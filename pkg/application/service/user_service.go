package service

import (
	"fmt"

	"github.com/ganyariya/itddd_go/pkg/boundary/dto"
	"github.com/ganyariya/itddd_go/pkg/boundary/irepository"
	"github.com/ganyariya/itddd_go/pkg/domain/entity"
	"github.com/ganyariya/itddd_go/pkg/domain/service"
	"github.com/ganyariya/itddd_go/pkg/domain/value"
)

/*
ドメインサービスはドメインを表現するために必要なもの。
それらドメインで表現された世界の問題を解決するものがアプリケーションサービス。
*/
type UserApplicationService struct {
	userRepository irepository.IUserRepository
	userService    service.UserService
}

func NewUserApplicationService(userRepo irepository.IUserRepository, userService service.UserService) *UserApplicationService {
	return &UserApplicationService{userRepository: userRepo, userService: userService}
}

func (uas *UserApplicationService) Register(userIdStr string, name string) (*dto.UserDTO, error) {
	userId, err := value.NewUserIdByIdString(userIdStr)
	if err != nil {
		return nil, err
	}
	user := entity.NewUser(userId, name)
	if uas.userService.Exists(user) {
		return nil, fmt.Errorf("exists")
	}

	uas.userRepository.Save(user)
	return dto.NewUserDTOByUser(user), nil
}

func (uas *UserApplicationService) Get(userIdStr string) (*dto.UserDTO, error) {
	userId, err := value.NewUserIdByIdString(userIdStr)
	if err != nil {
		return nil, err
	}

	user, err := uas.userRepository.Find(userId)
	if err != nil {
		return nil, err
	}
	return dto.NewUserDTOByUser(user), nil
}
