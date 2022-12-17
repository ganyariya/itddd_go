package interactor_test

import (
	"testing"

	"github.com/ganyariya/itddd_go/pkg/boundary/input"
	"github.com/ganyariya/itddd_go/pkg/domain/entity"
	"github.com/ganyariya/itddd_go/pkg/domain/interactor"
	"github.com/ganyariya/itddd_go/pkg/domain/value"
	"github.com/ganyariya/itddd_go/pkg/infrastructure/stub"
	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
	rawUserId := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	stubRepo := stub.NewUserStubRepository([]*entity.User{})
	interactor := interactor.NewUserCreateInteractor(
		stubRepo,
	)
	inputData := input.NewUserCreateInputData(rawUserId, "ganyariya")

	outputData, err := interactor.Handle(inputData)
	assert.NoError(t, err)
	assert.Equal(t, "ganyariya", outputData.User.Name)

	userId, err := value.NewUserId(rawUserId)
	assert.NoError(t, err)
	founded, err := stubRepo.Find(userId)
	assert.NoError(t, err)
	assert.Equal(t, "ganyariya", founded.Name)
}
