package specification

import (
	"github.com/ganyariya/itddd_go/pkg/domain/collection"
)

type CircleFullSpecification struct {
}

func NewCircleFullSpecification() *CircleFullSpecification {
	return &CircleFullSpecification{}
}

/*
複雑な仕様自体を Entity として切り出す

Circle Entity 内で JoinUser(user User, spec CircleFullSpecification) のようにしたいが循環 import になってしまう
サークルが満員か判定する仕様
*/
func (cfs *CircleFullSpecification) IsSatisfiedBy(circleMembers *collection.CircleMembers) (bool, error) {
	premiumCount := circleMembers.CountPremiumMembers()

	upper := 30
	if premiumCount >= 10 {
		upper = 50
	}
	return circleMembers.CountMembers() >= upper, nil
}
