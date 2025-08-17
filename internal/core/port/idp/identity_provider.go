package idp

import (
	"app/internal/core/shared_kernel/domain"
	"context"
)

type IdentityProvider interface {
	CreateUser(ctx context.Context, id domain.UserID, username string, email string, password string) (*domain.IdPUserId, error)
	ConfirmUser(ctx context.Context, id domain.IdPUserId) error
}
