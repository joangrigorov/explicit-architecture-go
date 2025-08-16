package usecases

import (
	. "app/internal/core/component/user/domain"
	. "app/internal/core/shared_kernel/events"
	"app/mock/core/component/user/application/repositories"
	"app/mock/core/port/events"
	"app/mock/core/port/uuid"
	"context"
	"errors"
	"testing"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRegistration_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	userRepository := repositories.NewMockUserRepository(ctrl)
	uuidGenerator := uuid.NewMockGenerator(ctrl)
	eventBus := events.NewMockEventBus(ctrl)

	useCase := NewRegistration(userRepository, uuidGenerator, eventBus)

	f := faker.New()

	username := f.Internet().User()
	password := f.Internet().Password()
	email := f.Internet().Email()
	fName := f.Person().FirstName()
	lName := f.Person().LastName()

	t.Run("success", func(t *testing.T) {
		id := f.UUID().V4()
		uuidGenerator.EXPECT().Generate().Return(id)
		userRepository.EXPECT().Create(ctx, gomock.AssignableToTypeOf(&User{})).Return(nil)
		eventBus.EXPECT().Publish(gomock.AssignableToTypeOf(UserCreated{}))

		assert.Nil(t, useCase.Execute(ctx, username, password, email, fName, lName))
	})

	t.Run("persistence error", func(t *testing.T) {
		id := f.UUID().V4()
		uuidGenerator.EXPECT().Generate().Return(id)
		userRepository.EXPECT().
			Create(ctx, gomock.AssignableToTypeOf(&User{})).
			Return(errors.New("persistence error"))

		assert.Equal(t, errors.New("persistence error"), useCase.Execute(ctx, username, password, email, fName, lName))
	})
}
