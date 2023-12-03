package inmemory

import (
	"bank/internal/domain"
	"sync"

	"github.com/gofrs/uuid"
)

type CardRepository struct {
	cards map[uuid.UUID]*domain.Card
	mutex sync.Mutex
}
