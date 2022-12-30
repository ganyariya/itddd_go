package factory

import "github.com/ganyariya/itddd_go/pkg/domain/value"

type ICircleFactory interface {
	Create(value.CircleName, value.UserId)
}
