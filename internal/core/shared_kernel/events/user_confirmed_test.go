package events

import (
	"app/internal/core/shared_kernel/domain"
	"testing"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

func TestUserConfirmed_ID(t *testing.T) {
	f := faker.New()
	userId := domain.UserID(f.UUID().V4())
	e := NewUserConfirmed(userId)

	assert.Equal(t, "app/internal/core/shared_kernel/events.UserConfirmed", e.ID().String())
}
