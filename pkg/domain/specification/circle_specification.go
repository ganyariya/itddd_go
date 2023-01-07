package specification

import (
	"github.com/ganyariya/itddd_go/pkg/boundary/irepository"
	"github.com/ganyariya/itddd_go/pkg/domain/entity"
)

type CircleFullSpecification struct {
	userRepository irepository.IUserRepository
}

func NewCircleFullSpecification(userRepository irepository.IUserRepository) *CircleFullSpecification {
	return &CircleFullSpecification{userRepository: userRepository}
}

/*
複雑な仕様自体を Entity として切り出す

Circle Entity 内で JoinUser(user User, spec CircleFullSpecification) のようにしたいが循環 import になってしまう
サークルが満員か判定する仕様
*/
func (cfs *CircleFullSpecification) IsSatisfiedBy(circle *entity.Circle) (bool, error) {
	users, err := cfs.userRepository.FindByIds(circle.MemberIds)
	if err != nil {
		return false, err
	}

	premiumCount := 0
	for _, u := range users {
		if u.IsPremium {
			premiumCount++
		}
	}

	upper := 30
	if premiumCount >= 10 {
		upper = 50
	}
	return circle.IsFull(upper), nil
}
