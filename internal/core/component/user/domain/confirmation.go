package domain

import (
	"app/internal/core/shared_kernel/domain"
	"time"
)

type ConfirmationID string

func (id ConfirmationID) String() string {
	return string(id)
}

type Confirmation struct {
	ID         ConfirmationID
	UserID     domain.UserID
	HMACSecret string
	CreatedAt  time.Time
}

func NewConfirmation(ID ConfirmationID, userID domain.UserID, hmacSecret string) *Confirmation {
	return &Confirmation{ID: ID, UserID: userID, HMACSecret: hmacSecret, CreatedAt: time.Now()}
}

func ReconstituteConfirmation(
	id ConfirmationID,
	userID domain.UserID,
	hmacSecret string,
	createdAt time.Time,
) *Confirmation {
	return &Confirmation{
		ID:         id,
		UserID:     userID,
		HMACSecret: hmacSecret,
		CreatedAt:  createdAt,
	}
}
