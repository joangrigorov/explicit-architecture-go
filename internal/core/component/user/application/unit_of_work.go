package application

import (
	"app/internal/core/component/user/application/repositories"
	eventBus "app/internal/core/port/events"
	"context"
)

// UnitOfWork used to wrap transactional business logic
type UnitOfWork interface {
	Do(context.Context, func(tx UnitOfWorkTx) error) error
}

type UnitOfWorkTx interface {
	EventBus() eventBus.EventBus
	UserRepository() repositories.UserRepository
}
