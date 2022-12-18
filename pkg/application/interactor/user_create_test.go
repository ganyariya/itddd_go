package interactor_test

import (
	"testing"

	"github.com/ganyariya/itddd_go/pkg/application/interactor"
	"github.com/ganyariya/itddd_go/pkg/boundary/input"
	"github.com/ganyariya/itddd_go/pkg/domain/entity"
	"github.com/ganyariya/itddd_go/pkg/domain/service"
	"github.com/ganyariya/itddd_go/pkg/infrastructure/stub"
	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
	stubRepo := stub.NewUserStubRepository([]*entity.User{})
	service := service.NewUserService(stubRepo)
	interactor := interactor.NewUserCreateInteractor(stubRepo, service)
	inputData := input.NewUserCreateInputData("ganyariya")

	outputData, err := interactor.Handle(inputData)
	assert.NoError(t, err)
	assert.Equal(t, "ganyariya", outputData.User.Name)

	userId := outputData.User.Id
	founded, err := stubRepo.Find(userId)
	assert.NoError(t, err)
	assert.Equal(t, "ganyariya", founded.Name)
}
