package usecases

import (
	"app/internal/core/component/user/application/errors"
	. "app/internal/core/component/user/application/repositories"
	. "app/internal/core/component/user/domain"
	. "app/internal/core/shared_kernel/domain"
	"context"
)

type SetIdPUserId struct {
	userRepository UserRepository
}

func NewSetIdPUserId(userRepository UserRepository) *SetIdPUserId {
	return &SetIdPUserId{userRepository: userRepository}
}

func (u *SetIdPUserId) Execute(
	ctx context.Context,
	userId string,
	idPUserId string,
) error {
	user, err := u.userRepository.GetById(ctx, UserID(userId))

	if err != nil {
		return errors.NewUserNotFoundError()
	}

	id := IdPUserId(idPUserId)
	user.IdPUserId = &id

	return u.userRepository.Update(ctx, user)
}
