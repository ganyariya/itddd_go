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

func TestUserService(t *testing.T) {
	tests := []struct {
		Name       string
		ChangeName string
	}{
		{Name: "ganyariya", ChangeName: "cake"},
		{Name: "ganyariya2", ChangeName: "cake2"},
	}

	userRepository := stub.NewUserStubRepository()
	userApplicationService := service.NewUserApplicationService(
		userRepository,
		dService.NewUserService(userRepository),
		factory.NewUserFactory(),
	)
	for _, tt := range tests {
		user, err := userApplicationService.Register(command.NewUserRegisterCommand(tt.Name, false))
		assert.NoError(t, err)
		assert.Equal(t, tt.Name, user.Name)

		user, err = userApplicationService.Get(command.NewUserGetCommand(user.Id))
		assert.NoError(t, err)
		assert.Equal(t, tt.Name, user.Name)

		user, err = userApplicationService.Update(command.NewUserUpdateCommand(user.Id, &tt.ChangeName))
		assert.NoError(t, err)
		assert.Equal(t, tt.ChangeName, user.Name)

		err = userApplicationService.Delete(command.NewUserDeleteCommand(user.Id))
		assert.NoError(t, err)

		// 削除済み
		_, err = userApplicationService.Get(command.NewUserGetCommand(user.Id))
		assert.Error(t, err)
	}
}
