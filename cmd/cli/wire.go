//go:build wireinject

package main

import (
	aService "github.com/ganyariya/itddd_go/pkg/application/service"
	"github.com/ganyariya/itddd_go/pkg/domain/factory"
	dService "github.com/ganyariya/itddd_go/pkg/domain/service"
	"github.com/ganyariya/itddd_go/pkg/infrastructure/stub"
	"github.com/google/wire"
)

func Initialize() *aService.UserApplicationService {
	wire.Build(
		aService.NewUserApplicationService,
		dService.NewUserService,
		stub.NewUserStubRepository,
		factory.NewUserFactory,
	)
	return &aService.UserApplicationService{}
}
