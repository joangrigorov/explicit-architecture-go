package handlers

import (
	. "app/internal/core/component/user/application/commands"
	port "app/internal/core/port/commands"
	"app/internal/core/port/uuid"
	"app/internal/infrastructure/commands"
	"app/internal/infrastructure/events"
	ent "app/internal/infrastructure/persistence/ent/generated/user"
	"app/internal/infrastructure/persistence/ent/user"
	"context"
)

func HandleRegisterUserCommand(
	userRepository *user.Repository,
	eventBus *events.SimpleEventBus,
	uuidGenerator uuid.Generator,
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

		bus := events.NewTransactionalEventBus(eventBus)

		defer func() {
			if err == nil {
				err = bus.Flush()
			} else {
				bus.Reset()
			}
		}()

		handler := NewRegisterUserCommandHandler(userRepository.WithTx(tx), eventBus, uuidGenerator)

		err = handler.Handle(ctx, command.(RegisterUserCommand))
		return err
	}
}
