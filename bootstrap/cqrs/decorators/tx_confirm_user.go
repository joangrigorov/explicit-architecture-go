package decorators

import (
	. "app/internal/core/component/user/application/commands"
	port "app/internal/core/port/cqrs"
	"app/internal/core/port/idp"
	"app/internal/infrastructure/cqrs/commands"
	ent "app/internal/infrastructure/persistence/ent/generated/user"
	"app/internal/infrastructure/persistence/ent/user"
	"context"
)

// TransactionalConfirmUserCommand runs the ConfirmUserCommandHandler in an Ent transaction,
// and it handles commit and rollbacks.
func TransactionalConfirmUserCommand(
	userRepository *user.Repository,
	idp idp.IdentityProvider,
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

		handler := NewConfirmUserCommandHandler(userRepository.WithTx(tx), idp)

		err = handler.Handle(ctx, command.(ConfirmUserCommand))
		return err
	}
}
