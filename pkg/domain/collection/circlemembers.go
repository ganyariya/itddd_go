package collection

import (
	"github.com/ganyariya/itddd_go/pkg/domain/entity"
	"github.com/ganyariya/itddd_go/pkg/domain/value"
)

/*
First-Class-Collection

配列や List などの「コレクション」に対する操作を Entity として閉じ込める

CircleMembers だけ Specification のために浮き出ている実装になっているため　本来は Entity とうまく整合性を取るべき
が、あまりに Collection に移しすぎると `Repository` の詰め替えが多く発生しそう
*/
type CircleMembers struct {
	CircleId value.CircleId
	Owner    *entity.User
	Members  []*entity.User
}

func NewCircleMembers(CircleId value.CircleId, Owner *entity.User, Members []*entity.User) *CircleMembers {
	return &CircleMembers{CircleId: CircleId, Owner: Owner, Members: Members}
}

func (cm *CircleMembers) CountMembers() int {
	return len(cm.Members) + 1
}

func (cm *CircleMembers) CountPremiumMembers() int {
	cnt := 0
	if cm.Owner.IsPremium {
		cnt++
	}
	for _, m := range cm.Members {
		if m.IsPremium {
			cnt++
		}
	}
	return cnt
}
