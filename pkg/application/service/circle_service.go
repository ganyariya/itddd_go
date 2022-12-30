package service

import (
	"fmt"

	"github.com/ganyariya/itddd_go/pkg/application/command"
	"github.com/ganyariya/itddd_go/pkg/boundary/dto"
	"github.com/ganyariya/itddd_go/pkg/boundary/irepository"
	"github.com/ganyariya/itddd_go/pkg/domain/factory"
	"github.com/ganyariya/itddd_go/pkg/domain/service"
	"github.com/ganyariya/itddd_go/pkg/domain/value"
)

type CircleApplicationService struct {
	circleRepository irepository.ICircleRepository
	userRepository   irepository.IUserRepository
	circleService    *service.CircleService
	circleFactory    factory.ICircleFactory
}

func NewCircleApplicationService(
	circleRepository irepository.ICircleRepository,
	userRepository irepository.IUserRepository,
	circleService *service.CircleService,
	circleFactory factory.ICircleFactory,
) *CircleApplicationService {
	return &CircleApplicationService{
		circleRepository: circleRepository,
		userRepository:   userRepository,
		circleService:    circleService,
		circleFactory:    circleFactory,
	}
}

func (cas *CircleApplicationService) Register(command command.CircleRegisterCommand) (*dto.CircleDTO, error) {
	userId, err := value.NewUserIdByIdString(command.UserId)
	if err != nil {
		return nil, err
	}
	user, err := cas.userRepository.Find(userId)
	if err != nil {
		return nil, err
	}

	circle, err := cas.circleFactory.Create(command.CircleName, user.Id)
	if err != nil {
		return nil, err
	}
	if cas.circleService.Exists(circle) {
		return nil, fmt.Errorf("circle already exists")
	}

	cas.circleRepository.Save(circle)
	return dto.NewCircleDTOByCircle(circle), nil
}

func (cas *CircleApplicationService) Join(command *command.CircleJoinCommand) (*dto.CircleDTO, error) {
	userId, err := value.NewUserIdByIdString(command.UserId)
	if err != nil {
		return nil, err
	}
	_, err = cas.userRepository.Find(userId)
	if err != nil {
		return nil, err
	}

	circleId, err := value.NewCircleIdByString(command.CircleId)
	if err != nil {
		return nil, err
	}
	circle, err := cas.circleRepository.FindByCircleId(circleId)
	if err != nil {
		return nil, err
	}

	circle.JoinUser(userId)

	cas.circleRepository.Save(circle)
	return dto.NewCircleDTOByCircle(circle), nil
}
