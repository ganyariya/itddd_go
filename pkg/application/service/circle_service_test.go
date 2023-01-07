package service_test

import (
	"testing"

	"github.com/ganyariya/itddd_go/pkg/application/command"
	"github.com/ganyariya/itddd_go/pkg/application/service"
	"github.com/ganyariya/itddd_go/pkg/domain/factory"
	dService "github.com/ganyariya/itddd_go/pkg/domain/service"
	"github.com/ganyariya/itddd_go/pkg/infrastructure/stub"
	"github.com/stretchr/testify/assert"
)

func TestCircleService(t *testing.T) {
	tests := []struct {
		UserName     string
		CircleName   string
		JoinUserName string
	}{
		{UserName: "ganyariya", CircleName: "tama", JoinUserName: "himari"},
	}

	userRepository := stub.NewUserStubRepository()
	circleRepository := stub.NewCircleRepository()
	circleApplicationService := service.NewCircleApplicationService(
		circleRepository,
		userRepository,
		dService.NewCircleService(circleRepository),
		factory.NewCircleFactory(),
	)
	userApplicationService := service.NewUserApplicationService(
		userRepository,
		dService.NewUserService(userRepository),
		factory.NewUserFactory(),
	)

	for _, tt := range tests {
		user, err := userApplicationService.Register(command.NewUserRegisterCommand(tt.UserName, false))
		assert.NoError(t, err)

		circle, err := circleApplicationService.Register(*command.NewCircleRegisterCommand(user.Id, tt.CircleName))
		assert.NoError(t, err)
		assert.Equal(t, tt.CircleName, circle.Name)
		assert.Equal(t, user.Id, circle.OwnerId)

		joinUser, err := userApplicationService.Register(command.NewUserRegisterCommand(tt.JoinUserName, false))
		assert.NoError(t, err)

		circle, err = circleApplicationService.Join(command.NewCircleJoinCommand(joinUser.Id, circle.Id))
		assert.NoError(t, err)
		assert.Contains(t, circle.MemberIds, joinUser.Id)
	}
}
