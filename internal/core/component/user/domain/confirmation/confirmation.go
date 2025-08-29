package confirmation

import (
	"app/internal/core/component/user/domain/user"
	"time"
)

type ID string

func (id ID) String() string {
	return string(id)
}

type Confirmation struct {
	ID         ID
	UserID     user.ID
	HMACSecret string
	CreatedAt  time.Time

	events []Event
}

func NewConfirmation(ID ID, userID user.ID, hmacSecret string) *Confirmation {
	return &Confirmation{
		ID:         ID,
		UserID:     userID,
		HMACSecret: hmacSecret,
		CreatedAt:  time.Now(),

		events: []Event{
			NewCreatedEvent(ID),
		},
	}
}

func ReconstituteConfirmation(
	id ID,
	userID user.ID,
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
