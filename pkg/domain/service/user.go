package service

import (
	"github.com/ganyariya/itddd_go/pkg/boundary/irepository"
	"github.com/ganyariya/itddd_go/pkg/domain/entity"
)

type UserService struct {
	userRepository irepository.IUserRepository
}

func NewUserService(userRepo irepository.IUserRepository) *UserService {
	return &UserService{userRepository: userRepo}
}

func (u *UserService) Exists(user *entity.User) bool {
	ret, _ := u.userRepository.Find(user.Id)
	return ret != nil
}
