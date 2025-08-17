package usecases

import (
	"app/internal/core/component/user/application"
	"app/internal/core/component/user/application/errors"
	"app/internal/core/port/idp"
	"app/internal/core/shared_kernel/domain"
	"app/internal/core/shared_kernel/events"
	"context"
)

type Confirmation struct {
	uow application.UnitOfWork
	idp idp.IdentityProvider
}

func NewConfirmation(uow application.UnitOfWork, idp idp.IdentityProvider) *Confirmation {
	return &Confirmation{uow: uow, idp: idp}
}

func (uc *Confirmation) Execute(ctx context.Context, userID string) error {
	return uc.uow.Do(ctx, func(tx application.UnitOfWorkTx) error {
		user, err := tx.UserRepository().GetById(ctx, domain.UserID(userID))

		if err != nil {
			return errors.NewUserNotFoundError()
		}

		if user.IdPUserId == nil {
			return errors.NewIdPUserNotConnectedError()
		}

		err = uc.idp.ConfirmUser(ctx, *user.IdPUserId)

		if err != nil {
			return errors.NewIdPRequestError(err)
		}

		user.Confirm()

		err = tx.UserRepository().Update(ctx, user)

		if err != nil {
			return err
		}

		return tx.EventBus().Publish(events.NewUserConfirmed(user.ID))
	})
}
