package commands

import (
	"app/internal/core/port/commands"
	"context"
	"errors"
	"fmt"
	"reflect"
)

type Next func(context.Context, commands.Command) error

type Middleware func(context.Context, commands.Command, Next) error

type SimpleCommandBus struct {
	handlers    map[reflect.Type][]Middleware
	middlewares []Middleware
}

func NewCommandBus(bus *SimpleCommandBus) commands.CommandBus {
	return bus
}

func NewSimpleCommandBus() *SimpleCommandBus {
	return &SimpleCommandBus{
		handlers: make(map[reflect.Type][]Middleware),
	}
}

func (s *SimpleCommandBus) Dispatch(ctx context.Context, command commands.Command) error {
	t := reflect.TypeOf(command)
	chain, ok := s.handlers[t]
	if !ok {
		return errors.New(fmt.Sprintf("no handlers found for command %T", command))
	}

	// prepend global middlewares
	chain = append(s.middlewares, chain...)

	// build the chain runner
	i := 0
	var exec Next
	exec = func(ctx context.Context, c commands.Command) error {
		if i >= len(chain) {
			return nil
		}
		current := chain[i]
		i++
		return current(ctx, c, exec)
	}

	return exec(ctx, command)
}

func Register[C commands.Command](bus *SimpleCommandBus, handlers ...Middleware) {
	t := typeOf[C]()
	bus.handlers[t] = append(bus.handlers[t], handlers...)
}

func Use(bus *SimpleCommandBus, handlers ...Middleware) {
	bus.middlewares = append(bus.middlewares, handlers...)
}

func typeOf[C any]() reflect.Type {
	var zero *C
	return reflect.TypeOf(zero).Elem()
}
