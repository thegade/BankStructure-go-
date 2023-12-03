package domain

import "github.com/gofrs/uuid"

type Bill struct {
	id       uuid.UUID
	balance  int
	limit    int
	isClosed bool
	userID   uuid.UUID
}

func (b *Bill) ID() uuid.UUID     { return b.id }
func (b *Bill) Balance() int      { return b.balance }
func (b *Bill) Limit() int        { return b.limit }
func (b *Bill) IsClosed() bool    { return b.isClosed }
func (b *Bill) UserID() uuid.UUID { return b.userID }

func NewBill(userID uuid.UUID) *Bill {
	return &Bill{
		id:     uuid.Must(uuid.NewV7()),
		userID: userID,
	}
}

func (b *Bill) Close() {
	b.isClosed = true
}
