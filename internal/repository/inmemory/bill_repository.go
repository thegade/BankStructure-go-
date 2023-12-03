package inmemory

import (
	"bank/internal/domain"
	"sync"

	"github.com/gofrs/uuid"
)

type BillRepository struct {
	bills map[uuid.UUID]*domain.Bill
	mutex sync.Mutex
}
