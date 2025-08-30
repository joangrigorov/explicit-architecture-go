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
	ID              ID
	UserID          user.ID
	UserEmailMasked string
	CSRFToken       CSRFToken
	ExpiresAt       time.Time
	UsedAt          *time.Time
	CreatedAt       time.Time

	events []Event
}

func (u *Verification) recordEvent(event Event) {
	u.events = append(u.events, event)
}

func (u *Verification) ResetEvents() {
	u.events = make([]Event, 0)
}

func (u *Verification) Events() []Event {
	return u.events
}

func NewVerification(ID ID, userID user.ID, userEmailMasked string, csrfToken CSRFToken) *Verification {
	return &Verification{
		ID:              ID,
		UserID:          userID,
		UserEmailMasked: userEmailMasked,
		CSRFToken:       csrfToken,
		ExpiresAt:       time.Now().Add(time.Hour * 24 * 2), // two days
		CreatedAt:       time.Now(),

		events: []Event{
			NewCreatedEvent(ID),
		},
	}
}
