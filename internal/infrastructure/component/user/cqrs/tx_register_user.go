package cqrs

import (
	"app/internal/core/component/user/application/commands/register_user"
	port "app/internal/core/port/cqrs"
	"app/internal/core/port/errors"
	usrAdapter "app/internal/infrastructure/component/user/persistence/ent"
	userEnt "app/internal/infrastructure/component/user/persistence/ent/generated"
	"app/internal/infrastructure/framework/cqrs/commands"
	"app/internal/infrastructure/framework/events"
	"app/internal/infrastructure/framework/persistence/pgsql"
	"context"
)

// TransactionalRegisterUserCommand runs the Handler in
// - an Ent transaction, handling commits and rollbacks.
// - uses TransactionalEventBus which flushes collected events only after successful command handling.
type TransactionalRegisterUserCommand struct {
	userRepository *usrAdapter.UserRepository
	eventBus       *events.SimpleEventBus
	entClient      *userEnt.Client
	errors         errors.ErrorFactory
}

func NewTransactionalRegisterUserCommand(
	userRepository *usrAdapter.UserRepository,
	eventBus *events.SimpleEventBus,
	entClient *userEnt.Client,
	errors errors.ErrorFactory,
) *TransactionalRegisterUserCommand {
	return &TransactionalRegisterUserCommand{
		userRepository: userRepository,
		eventBus:       eventBus,
		entClient:      entClient,
		errors:         errors,
	}
}

func (t *TransactionalRegisterUserCommand) Provide(ctx context.Context, command port.Command, _ commands.Next) (err error) {
	txEventBus := events.NewTransactionalEventBus(t.eventBus)
	defer events.CloseEventBus(ctx, txEventBus, &err)

	tx, err := t.entClient.Tx(ctx)
	if err != nil {
		return err
	}
	defer pgsql.CloseTx(tx, &err)

	userRepository := t.userRepository.
		WithTx(tx).
		WithEventBus(txEventBus)

	handler := register_user.NewHandler(userRepository, t.errors)
	return handler.Handle(ctx, command.(register_user.Command))
}
