package service

import (
	"github.com/ganyariya/itddd_go/pkg/boundary/irepository"
	"github.com/ganyariya/itddd_go/pkg/domain/entity"
)

/*
ドメイン内で Entity 単体にもたせられない処理を ドメインサービスが担当する
ドメインの世界と知識を表現するものであり　解決するものがアプリケーションサービス
*/
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
