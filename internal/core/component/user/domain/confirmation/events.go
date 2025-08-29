package confirmation

import (
	"app/internal/core/shared_kernel/domain"
	"time"
)

type Event interface {
	ConfirmationID() ID
	CreatedAt() time.Time
}

type CreatedEvent struct {
	confirmationID ID

	domain.WithCreatedAt
}

func NewCreatedEvent(confirmationID ID) *CreatedEvent {
	return &CreatedEvent{confirmationID: confirmationID, WithCreatedAt: domain.NewWithCreatedAtNow()}
}

func (c CreatedEvent) ConfirmationID() ID {
	return c.confirmationID
}
