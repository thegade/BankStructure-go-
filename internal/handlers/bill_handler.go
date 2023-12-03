package handlers

import (
	"encoding/json"
	"github.com/gofrs/uuid"
)

type GETBillRequest struct {
	ID uuid.UUID
}

type GETBillResponse struct {
	id       uuid.UUID
	isClosed bool
	userID   uuid.UUID
	limit    int
	balance  int
}

type POSTBillRequest struct {
	ID uuid.UUID
}

type POSTBillResponse struct {
	id       uuid.UUID
	isClosed bool
	userID   uuid.UUID
}

func (response *POSTBillResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID       uuid.UUID
		IsClosed bool
	}{
		ID:       response.id,
		IsClosed: response.isClosed,
	})
}

func (response *GETBillResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID       uuid.UUID
		IsOpened bool
		UserID   uuid.UUID
		Limit    int
		Balance  int
	}{
		ID:       response.id,
		IsOpened: response.isClosed,
		UserID:   response.userID,
		Limit:    response.limit,
		Balance:  response.balance,
	})
}
