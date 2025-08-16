package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserNotFoundError_Error(t *testing.T) {
	err := NewUserNotFoundError()

	assert.Equal(t, &UserNotFoundError{}, err)
	assert.Equal(t, "user not found", err.Error())
}
