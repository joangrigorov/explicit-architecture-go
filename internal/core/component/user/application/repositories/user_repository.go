package repositories

import (
	"app/internal/core/component/user/domain"
	sk "app/internal/core/shared_kernel/domain"
	"context"
)

type UserRepository interface {
	GetById(context.Context, sk.UserID) (*domain.User, error)
	GetByIdPUserId(context.Context, sk.IdPUserID) (*domain.User, error)
	Create(context.Context, *domain.User) error
	Update(ctx context.Context, user *domain.User) error
}
