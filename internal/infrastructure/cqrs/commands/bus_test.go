package commands

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

func TestSimpleCommandBus_Dispatch(t *testing.T) {
	ctx := context.Background()
	bus := NewSimpleCommandBus()

	Register[myCmd](bus, func(ctx context.Context, command cqrs.Command, next Next) error {
		return errors.New("this is an expected error")
	})

	err := bus.Dispatch(ctx, myCmd{})

	assert.Equal(t, errors.New("this is an expected error"), err)
}
