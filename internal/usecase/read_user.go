package usecase

import (
	"bank/internal/domain"
)

type ReadUserUseCase struct {
	userRepository domain.UserRepository
}

func NewReadUserUseCase(userRepository domain.UserRepository) *ReadUserUseCase {
	return &ReadUserUseCase{
		userRepository: userRepository,
	}
}

type ReadUserCommand struct {
	Login    string
	Password string
}

func (useCase *ReadUserUseCase) Handle(command *ReadUserCommand) (*domain.User, error) {
	user, err := useCase.userRepository.FindUser(command.Login, command.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
