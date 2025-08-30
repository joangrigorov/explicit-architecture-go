package verification

import (
	"app/internal/core/shared_kernel/domain/mixin"
	"time"
)

type Event interface {
	VerificationID() ID
	CreatedAt() time.Time
}

type CreatedEvent struct {
	verificationID ID

	mixin.WithCreatedAt
}

func NewCreatedEvent(verificationID ID) *CreatedEvent {
	return &CreatedEvent{verificationID: verificationID, WithCreatedAt: mixin.NewWithCreatedAtNow()}
}

func (c CreatedEvent) VerificationID() ID {
	return c.verificationID
}
