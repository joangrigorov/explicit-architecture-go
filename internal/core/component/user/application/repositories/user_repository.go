package repositories

import (
	"app/internal/core/component/user/domain"
	"context"
)

type UserRepository interface {
	GetByIdPUserId(context.Context, domain.IdPUserId) (*domain.User, error)
	Create(context.Context, *domain.User) error
}
