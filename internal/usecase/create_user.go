package usecase

import (
	"bank/internal/domain"

	"github.com/gofrs/uuid"
)

type CreateUserUseCase struct {
	userRepository domain.UserRepository
}

func NewCreateUserUseCase(userRepository domain.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: userRepository,
	}
}

type CreateUserCommand struct {
	ID       uuid.UUID
	Login    string
	Password string
	Fullname string
}

func (useCase *CreateUserUseCase) Handle(command *CreateUserCommand) (*domain.User, error) {
	user := domain.NewUser(command.ID, command.Login, command.Password, command.Fullname)
	_, err := useCase.userRepository.Save(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
