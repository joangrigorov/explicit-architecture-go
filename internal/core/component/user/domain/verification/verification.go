package verification

import (
	"app/internal/core/component/user/domain/user"
	"time"
)

type ID string

func (id ID) String() string {
	return string(id)
}

type Verification struct {
	ID        ID
	UserID    user.ID
	CSRFToken CSRFToken
	ExpiresAt time.Time
	UsedAt    *time.Time
	CreatedAt time.Time

	events []Event
}

func NewVerification(ID ID, userID user.ID, csrfToken CSRFToken) *Verification {
	return &Verification{
		ID:        ID,
		UserID:    userID,
		CSRFToken: csrfToken,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 2), // two days
		CreatedAt: time.Now(),

		events: []Event{
			NewCreatedEvent(ID),
		},
	}
}
