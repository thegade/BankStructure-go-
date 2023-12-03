package inmemory

import (
	"bank/internal/domain"
	"sync"

	"github.com/gofrs/uuid"
)

type UserRepository struct {
	users map[uuid.UUID]*domain.User
	mutex sync.Mutex
}

func (repository *UserRepository) Save(user *domain.User) error {
	repository.mutex.Lock()
	defer repository.mutex.Unlock()
	repository.users[user.ID()] = user
	return nil
}

func (repository *UserRepository) FindByID(id uuid.UUID) (*domain.User, error) {
	repository.mutex.Lock()
	defer repository.mutex.Unlock()
	user := repository.users[id]
	return user, nil
}
func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[uuid.UUID]*domain.User),
	}
}
