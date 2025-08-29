package idp

import (
	"app/internal/core/component/user/domain/user"
	"context"
)

type IdentityProvider interface {
	CreateUser(ctx context.Context, id user.ID, username string, email string, password string) (*user.IdPUserID, error)
	ConfirmUser(context.Context, user.IdPUserID) error
}
