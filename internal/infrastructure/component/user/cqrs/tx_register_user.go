package cqrs

import (
	. "app/internal/core/component/user/application/commands"
	port "app/internal/core/port/cqrs"
	"app/internal/core/port/errors"
	usrAdapter "app/internal/infrastructure/component/user/persistence/ent"
	userEnt "app/internal/infrastructure/component/user/persistence/ent/generated"
	"app/internal/infrastructure/framework/cqrs/commands"
	"app/internal/infrastructure/framework/event_bus"
	"app/internal/infrastructure/framework/persistence/pgsql"
	"context"
)

// TransactionalRegisterUserCommand runs the RegisterUserCommandHandler in
// - an Ent transaction, handling commits and rollbacks.
// - uses TransactionalEventBus which flushes collected events only after successful command handling.
type TransactionalRegisterUserCommand struct {
	userRepository *usrAdapter.UserRepository
	eventBus       *event_bus.SimpleEventBus
	entClient      *userEnt.Client
	errors         errors.ErrorFactory
}

func NewTransactionalRegisterUserCommand(
	userRepository *usrAdapter.UserRepository,
	eventBus *event_bus.SimpleEventBus,
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
	txEventBus := event_bus.NewTransactionalEventBus(t.eventBus)
	defer event_bus.CloseEventBus(ctx, txEventBus, &err)

	tx, err := t.entClient.Tx(ctx)
	if err != nil {
		return err
	}
	defer pgsql.CloseTx(tx, &err)

	handler := NewRegisterUserCommandHandler(t.userRepository.WithTx(tx), txEventBus, t.errors)
	return handler.Handle(ctx, command.(RegisterUserCommand))
}
