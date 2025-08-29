package repositories

import (
	"app/internal/core/component/user/domain/verification"
	"context"
)

type VerificationRepository interface {
	Create(context.Context, *verification.Verification) error
	GetByID(context.Context, verification.ID) (*verification.Verification, error)
	Expire(context.Context, *verification.Verification) error
}
