package usecases

import (
	"app/internal/core/component/user/application/errors"
	"app/internal/core/component/user/application/repositories"
	"app/internal/core/port/idp"
	"app/internal/core/shared_kernel/domain"
	"context"
	errors2 "errors"
)

type CreateIdPUser struct {
	userRepository repositories.UserRepository
	idp            idp.IdentityProvider
}

func NewCreateIdPUser(userRepository repositories.UserRepository, idp idp.IdentityProvider) *CreateIdPUser {
	return &CreateIdPUser{userRepository: userRepository, idp: idp}
}

func (u *CreateIdPUser) Execute(
	ctx context.Context,
	userID domain.UserID,
	username string,
	email string,
	password string,
) error {
	user, err := u.userRepository.GetById(ctx, userID)

	if err != nil {
		return errors2.New(userID.String())
	}

	idpUserId, err := u.idp.CreateUser(ctx, userID, username, email, password)

	if err != nil {
		return errors.NewCannotCreateIdPUserError(err)
	}

	user.IdPUserId = idpUserId

	return u.userRepository.Update(ctx, user)
}
