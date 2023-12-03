package domain

import "github.com/gofrs/uuid"

type Card struct {
	id       uuid.UUID
	isActive bool
	userID   uuid.UUID
}

func (c *Card) ID() uuid.UUID     { return c.id }
func (c *Card) IsActive() bool    { return c.isActive }
func (c *Card) UserID() uuid.UUID { return c.userID }

func NewCard(userID uuid.UUID) *Card {
	return &Card{
		id:     uuid.Must(uuid.NewV7()),
		userID: userID,
	}
}
func (c *Card) Block() {
	c.isActive = false
}
