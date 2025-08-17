package user

import (
	port "app/internal/core/component/user/application"
	"app/internal/core/component/user/application/repositories"
	eventBus "app/internal/core/port/events"
	"app/internal/infrastructure/events"
	"app/internal/infrastructure/persistence/ent/generated/user"
	"context"
)

type UnitOfWorkTx struct {
	tx             *user.Tx
	userRepository *Repository
	eventBus       *events.TransactionalEventBus
}

func (u *UnitOfWorkTx) EventBus() eventBus.EventBus {
	return u.eventBus
}

func (u *UnitOfWorkTx) UserRepository() repositories.UserRepository {
	return u.userRepository
}

func NewUnitOfWorkTx(tx *user.Tx, userRepository *Repository, eventBus *events.TransactionalEventBus) *UnitOfWorkTx {
	return &UnitOfWorkTx{
		tx:             tx,
		userRepository: userRepository.WithTx(tx),
		eventBus:       eventBus,
	}
}

type UnitOfWork struct {
	entClient      *user.Client
	userRepository *Repository
	eventBus       *events.SimpleEventBus
}

func NewUnitOfWork(
	userRepository *Repository,
	entClient *user.Client,
	eventBus *events.SimpleEventBus,
) port.UnitOfWork {
	return &UnitOfWork{
		userRepository: userRepository,
		entClient:      entClient,
		eventBus:       eventBus,
	}
}

func (u *UnitOfWork) Do(ctx context.Context, fn func(tx port.UnitOfWorkTx) error) error {
	tx, err := u.entClient.Tx(ctx)
	if err != nil {
		return err
	}

	bus := events.NewTransactionalEventBus(u.eventBus)
	txWrapper := NewUnitOfWorkTx(tx, u.userRepository, bus)

	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			panic(r)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	defer func() {
		if err == nil {
			err = bus.Flush()
		} else {
			bus.Reset()
		}
	}()

	err = fn(txWrapper)
	return err
}
