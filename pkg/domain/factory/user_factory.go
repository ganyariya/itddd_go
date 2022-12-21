package factory

import (
	"github.com/ganyariya/itddd_go/pkg/domain/entity"
	"github.com/ganyariya/itddd_go/pkg/domain/value"
)

type IUserFactory interface {
	Create(userName string) (*entity.User, error)
}

/*
User の生成を担当する
（採番処理などを行う場合など、難しい複雑な処理を行う場合
ファクトリにインスタンス生成を依頼する）

Domain の世界に閉じさせたいが `IUserRepository` を使うことになってしまう
Factory ならびに Domain Service は Boundary とのやり取りも行うことになりそう
*/
type UserFactory struct {
}

func NewUserFactory() IUserFactory {
	return &UserFactory{}
}

func (uf *UserFactory) Create(userName string) (*entity.User, error) {
	userId, err := value.NewUserId()
	if err != nil {
		return nil, err
	}
	user := entity.NewUser(userId, userName)
	return user, nil
}
