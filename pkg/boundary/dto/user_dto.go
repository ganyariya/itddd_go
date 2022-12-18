package dto

import "github.com/ganyariya/itddd_go/pkg/domain/entity"

/*
ドメイン Entity を直接外部に出してしまうと　クライアント側で操作される可能性がある
ABR の外側に出す場合は DTO に置き換える
*/
type UserDTO struct {
	Id   string
	Name string
}

func NewUserDTOByUser(user *entity.User) *UserDTO {
	return &UserDTO{
		Id:   user.Id.Id,
		Name: user.Name,
	}
}
