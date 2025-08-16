package usecases

import (
	. "app/internal/core/component/user/application/errors"
	. "app/internal/core/component/user/domain"
	"app/internal/core/shared_kernel/domain"
	"app/mock/core/component/user/application/repositories"
	"context"
	"errors"
	"testing"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSetIdPUserId_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	userRepository := repositories.NewMockUserRepository(ctrl)
	useCase := NewSetIdPUserId(userRepository)

	f := faker.New()
	userId := f.UUID().V4()
	idPUserId := f.UUID().V4()

	t.Run("success", func(t *testing.T) {
		user := &User{}
		userRepository.EXPECT().
			GetById(ctx, gomock.AssignableToTypeOf(domain.UserID(userId))).
			DoAndReturn(func(ctx context.Context, actualUserId domain.UserID) (*User, error) {
				assert.Equal(t, userId, string(actualUserId))

				return user, nil
			})

		userRepository.EXPECT().Update(ctx, user).Return(nil)

		assert.Nil(t, useCase.Execute(ctx, userId, idPUserId))
		assert.Equal(t, idPUserId, string(*user.IdPUserId))
	})

	t.Run("user not found", func(t *testing.T) {
		userRepository.EXPECT().
			GetById(ctx, gomock.AssignableToTypeOf(domain.UserID(userId))).
			Return(nil, errors.New("not found"))

		err := useCase.Execute(ctx, userId, idPUserId)

		assert.Equal(t, &UserNotFoundError{}, err)
	})
}
