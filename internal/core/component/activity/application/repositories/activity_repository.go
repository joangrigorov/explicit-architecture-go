package repositories

import (
	"app/internal/core/component/activity/domain"
	"context"
)

type ActivityRepository interface {
	GetById(context.Context, domain.ActivityId) (*domain.Activity, error)
	GetAll(context.Context) ([]*domain.Activity, error)
	Create(context.Context, *domain.Activity) error
	Update(context.Context, *domain.Activity) error
	Delete(context.Context, *domain.Activity) error
}
