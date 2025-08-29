package repositories

import (
	"app/internal/core/component/user/domain/user"
	"context"
)

type UserRepository interface {
	GetById(context.Context, user.ID) (*user.User, error)
	GetByEmail(context.Context, user.Email) (*user.User, error)
	GetByUsername(context.Context, user.Username) (*user.User, error)
	GetByIdPUserId(context.Context, user.IdPUserID) (*user.User, error)
	Create(context.Context, *user.User) error
	Update(ctx context.Context, user *user.User) error
}
