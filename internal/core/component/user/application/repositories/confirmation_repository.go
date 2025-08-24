package repositories

import (
	"app/internal/core/component/user/domain"
	"context"
)

type ConfirmationRepository interface {
	Create(context.Context, *domain.Confirmation) error
	GetByID(context.Context, domain.ConfirmationID) (*domain.Confirmation, error)
	Expire(context.Context, *domain.Confirmation) error
}
