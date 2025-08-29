package verification

import (
	"app/internal/core/shared_kernel/domain"
	"time"
)

type Event interface {
	VerificationID() ID
	CreatedAt() time.Time
}

type CreatedEvent struct {
	verificationID ID

	domain.WithCreatedAt
}

func NewCreatedEvent(verificationID ID) *CreatedEvent {
	return &CreatedEvent{verificationID: verificationID, WithCreatedAt: domain.NewWithCreatedAtNow()}
}

func (c CreatedEvent) VerificationID() ID {
	return c.verificationID
}
