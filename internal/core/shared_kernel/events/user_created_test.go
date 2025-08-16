package events

import (
	"app/internal/core/shared_kernel/domain"
	"testing"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

func TestUserCreated_ID(t *testing.T) {
	f := faker.New()
	userId := domain.UserID(f.UUID().V4())
	email := f.Internet().Email()
	password := f.Internet().Password()

	e := NewUserCreated(userId, email, password)

	assert.NotNil(t, e)
	assert.Equal(t, "app/internal/core/shared_kernel/events.UserCreated", e.ID().String())
}
