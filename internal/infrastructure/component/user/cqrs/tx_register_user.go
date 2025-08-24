package cqrs

import (
	. "app/internal/core/component/user/application/commands"
	port "app/internal/core/port/cqrs"
	usrAdapter "app/internal/infrastructure/component/user/persistence/ent"
	userEnt "app/internal/infrastructure/component/user/persistence/ent/generated"
	"app/internal/infrastructure/framework/cqrs/commands"
	"app/internal/infrastructure/framework/event_bus"
	"context"
)

// TransactionalRegisterUserCommand runs the RegisterUserCommandHandler in
// - an Ent transaction, handling commits and rollbacks.
// - uses TransactionalEventBus which flushes collected events only after successful command handling.
func TransactionalRegisterUserCommand(
	userRepository *usrAdapter.UserRepository,
	eventBus *event_bus.SimpleEventBus,
	entClient *userEnt.Client,
) commands.Middleware {
	return func(ctx context.Context, command port.Command, next commands.Next) error {
		tx, err := entClient.Tx(ctx)
		if err != nil {
			return err
		}

		txEventBus := event_bus.NewTransactionalEventBus(eventBus)

		defer func() {
			if err == nil {
				err = txEventBus.Flush()
			} else {
				txEventBus.Reset()
			}
		}()

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

		handler := NewRegisterUserCommandHandler(userRepository.WithTx(tx), txEventBus)

		err = handler.Handle(ctx, command.(RegisterUserCommand))
		return err
	}
}
