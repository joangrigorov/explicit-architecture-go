package repositories

import (
	"app/internal/core/component/user/domain/confirmation"
	"context"
)

type ConfirmationRepository interface {
	Create(context.Context, *confirmation.Confirmation) error
	GetByID(context.Context, confirmation.ID) (*confirmation.Confirmation, error)
	Expire(context.Context, *confirmation.Confirmation) error
}
