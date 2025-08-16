package usecases

import (
	. "app/internal/core/component/user/application/errors"
	. "app/internal/core/component/user/domain"
	"app/internal/core/shared_kernel/domain"
	. "app/internal/core/shared_kernel/events"
	"app/mock/core/component/user/application/repositories"
	"app/mock/core/port/events"
	"context"
	"errors"
	"testing"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestConfirmation_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	userRepository := repositories.NewMockUserRepository(ctrl)
	eventBus := events.NewMockEventBus(ctrl)

	useCase := NewConfirmation(userRepository, eventBus)

	f := faker.New()
	userId := f.UUID().V4()

	t.Run("success", func(t *testing.T) {
		user := &User{}
		userRepository.EXPECT().
			GetById(ctx, gomock.AssignableToTypeOf(domain.UserID(userId))).
			DoAndReturn(func(ctx context.Context, actualUserId domain.UserID) (*User, error) {
				assert.Equal(t, userId, string(actualUserId))

				return user, nil
			})
		userRepository.EXPECT().Update(ctx, user).Return(nil)
		eventBus.EXPECT().Publish(gomock.AssignableToTypeOf(UserConfirmed{}))

		assert.Nil(t, user.ConfirmedAt)
		assert.Nil(t, useCase.Execute(ctx, userId))
		assert.NotNil(t, user.ConfirmedAt)
	})

	t.Run("update error", func(t *testing.T) {
		user := &User{}
		userRepository.EXPECT().
			GetById(ctx, gomock.AssignableToTypeOf(domain.UserID(userId))).
			DoAndReturn(func(ctx context.Context, actualUserId domain.UserID) (*User, error) {
				assert.Equal(t, userId, string(actualUserId))

				return user, nil
			})
		userRepository.EXPECT().Update(ctx, user).Return(errors.New("persistence error"))

		assert.NotNil(t, useCase.Execute(ctx, userId))
	})

	t.Run("user not found", func(t *testing.T) {
		userRepository.EXPECT().
			GetById(ctx, gomock.AssignableToTypeOf(domain.UserID(userId))).
			Return(nil, errors.New("not found"))

		err := useCase.Execute(ctx, userId)

		assert.Equal(t, &UserNotFoundError{}, err)
	})
}
