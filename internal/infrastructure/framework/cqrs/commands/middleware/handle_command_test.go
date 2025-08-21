package middleware

import (
	"app/internal/core/port/cqrs"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type myCmd struct {
	someString string
}

type myCmdHandler struct{}

func (m *myCmdHandler) Handle(ctx context.Context, c myCmd) error {
	return errors.New("test error just to see if it works")
}

func TestHandleCommand(t *testing.T) {
	mw := HandleCommand[myCmd](&myCmdHandler{})
	ignoredNext := func(ctx context.Context, command cqrs.Command) error {
		return nil
	}

	err := mw(context.Background(), myCmd{}, ignoredNext)

	assert.Equal(t, errors.New("test error just to see if it works"), err)
}
