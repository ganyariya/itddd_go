package usecase

import "github.com/ganyariya/itddd_go/pkg/boundary/input"

type UserCreateUseCase interface {
	Handle(input *input.UserCreateInputData) (*input.UserCreateInputData, error)
}
