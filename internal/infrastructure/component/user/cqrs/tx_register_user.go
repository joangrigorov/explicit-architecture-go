package cqrs

import (
	. "app/internal/core/component/user/application/commands"
	port "app/internal/core/port/cqrs"
	ent2 "app/internal/infrastructure/component/user/persistence/ent"
	ent "app/internal/infrastructure/component/user/persistence/ent/generated"
	"app/internal/infrastructure/framework/cqrs/commands"
	event_bus2 "app/internal/infrastructure/framework/event_bus"
	"context"
)

// HandleRegisterUserCommand runs the RegisterUserCommandHandler in
// - an Ent transaction, handling commits and rollbacks.
// - uses TransactionalEventBus which flushes collected events only after successful command handling.
func HandleRegisterUserCommand(
	userRepository *ent2.Repository,
	eventBus *event_bus2.SimpleEventBus,
	entClient *ent.Client,
) commands.Middleware {
	return func(ctx context.Context, command port.Command, next commands.Next) error {
		tx, err := entClient.Tx(ctx)
		if err != nil {
			return err
		}

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

		bus := event_bus2.NewTransactionalEventBus(eventBus)

		defer func() {
			if err == nil {
				err = bus.Flush()
			} else {
				bus.Reset()
			}
		}()

		handler := NewRegisterUserCommandHandler(userRepository.WithTx(tx), eventBus)

		err = handler.Handle(ctx, command.(RegisterUserCommand))
		return err
	}
}
