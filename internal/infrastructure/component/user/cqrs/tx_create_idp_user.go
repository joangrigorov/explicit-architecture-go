package cqrs

import (
	. "app/internal/core/component/user/application/commands"
	port "app/internal/core/port/cqrs"
	"app/internal/core/port/idp"
	usrAdapter "app/internal/infrastructure/component/user/persistence/ent"
	userEnt "app/internal/infrastructure/component/user/persistence/ent/generated"
	"app/internal/infrastructure/framework/cqrs/commands"
	"app/internal/infrastructure/framework/event_bus"
	"app/internal/infrastructure/framework/persistence/pgsql"
	"context"
)

// TransactionalCreateIdPUserCommand runs the CreateIdPUserCommandHandler in an Ent transaction,
// - an Ent transaction, handling commits and rollbacks.
// - uses TransactionalEventBus which flushes collected events only after successful command handling.
type TransactionalCreateIdPUserCommand struct {
	userRepository *usrAdapter.UserRepository
	eventBus       *event_bus.SimpleEventBus
	idp            idp.IdentityProvider
	entClient      *userEnt.Client
}

func NewTransactionalCreateIdPUserCommand(
	userRepository *usrAdapter.UserRepository,
	eventBus *event_bus.SimpleEventBus,
	idp idp.IdentityProvider,
	entClient *userEnt.Client,
) *TransactionalCreateIdPUserCommand {
	return &TransactionalCreateIdPUserCommand{
		userRepository: userRepository,
		eventBus:       eventBus,
		idp:            idp,
		entClient:      entClient,
	}
}

func (t *TransactionalCreateIdPUserCommand) Provide(ctx context.Context, command port.Command, next commands.Next) (err error) {
	bus := event_bus.NewTransactionalEventBus(t.eventBus)
	defer event_bus.CloseEventBus(ctx, bus, &err)

	tx, err := t.entClient.Tx(ctx)
	if err != nil {
		return err
	}
	defer pgsql.CloseTx(tx, &err)

	handler := NewCreateIdPUserHandler(t.userRepository.WithTx(tx), t.idp, bus)
	return handler.Handle(ctx, command.(CreateIdPUserCommand))
}
