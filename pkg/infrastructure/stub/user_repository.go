package stub

import (
	"fmt"

	"github.com/ganyariya/itddd_go/pkg/domain/entity"
	"github.com/ganyariya/itddd_go/pkg/domain/value"
)

type UserStubRepository struct {
	users []*entity.User
}

func NewUserStubRepository(users []*entity.User) *UserStubRepository {
	return &UserStubRepository{users: users}
}

func (u *UserStubRepository) Save(user *entity.User) error {
	for i := range u.users {
		if u.users[i].Id == user.Id {
			u.users[i] = user
			return nil
		}
	}
	u.users = append(u.users, user)
	return nil
}
func (u *UserStubRepository) Find(userId value.UserId) (*entity.User, error) {
	for _, user := range u.users {
		if user.Id == userId {
			return user, nil
		}
	}
	return nil, fmt.Errorf("not found")
}
func (u *UserStubRepository) FindByName(name string) (*entity.User, error) {
	for _, user := range u.users {
		if user.Name == name {
			return user, nil
		}
	}
	return nil, fmt.Errorf("not found")
}
func (u *UserStubRepository) Delete(user *entity.User) {
	for i, uu := range u.users {
		if uu.Id == user.Id {
			u.users = append(u.users[:i], u.users[i+1:]...)
			return
		}
	}
}