package service_test

import (
	"testing"

	"github.com/ganyariya/itddd_go/pkg/application/service"
	"github.com/ganyariya/itddd_go/pkg/domain/entity"
	dService "github.com/ganyariya/itddd_go/pkg/domain/service"
	"github.com/ganyariya/itddd_go/pkg/domain/value"
	"github.com/ganyariya/itddd_go/pkg/infrastructure/stub"
	"github.com/stretchr/testify/assert"
)

func TestUserService(t *testing.T) {
	tests := []struct {
		Name string
	}{
		{Name: "ganyariya"},
		{Name: "ganyariya2"},
	}

	userRepository := stub.NewUserStubRepository([]*entity.User{})
	userApplicationService := service.NewUserApplicationService(
		userRepository,
		*dService.NewUserService(userRepository),
	)
	for _, tt := range tests {
		userId, err := value.NewUserId()
		assert.NoError(t, err)

		user, err := userApplicationService.Register(userId.Id, tt.Name)
		assert.NoError(t, err)
		assert.Equal(t, tt.Name, user.Name)
		assert.Equal(t, userId.Id, user.Id)

		user, err = userApplicationService.Get(userId.Id)
		assert.NoError(t, err)
		assert.Equal(t, tt.Name, user.Name)
		assert.Equal(t, userId.Id, user.Id)

		_, err = userApplicationService.Register(userId.Id, tt.Name)
		assert.Error(t, err)
	}
}
