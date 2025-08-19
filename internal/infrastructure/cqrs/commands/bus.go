package commands

import (
	"app/internal/core/port/cqrs"
	"context"
	"errors"
	"fmt"
	"reflect"
)

type Next func(context.Context, cqrs.Command) error

type Middleware func(context.Context, cqrs.Command, Next) error

type SimpleCommandBus struct {
	handlers    map[reflect.Type][]Middleware
	middlewares []Middleware
}

func NewCommandBus(bus *SimpleCommandBus) cqrs.CommandBus {
	return bus
}

func NewSimpleCommandBus() *SimpleCommandBus {
	return &SimpleCommandBus{
		handlers: make(map[reflect.Type][]Middleware),
	}
}

func (b *SimpleCommandBus) Dispatch(ctx context.Context, command cqrs.Command) error {
	t := reflect.TypeOf(command)
	chain, ok := b.handlers[t]
	if !ok {
		return errors.New(fmt.Sprintf("no handlers found for command %T", command))
	}

	// prepend global middlewares
	chain = append(b.middlewares, chain...)

	// build the chain runner
	i := 0
	var exec Next
	exec = func(ctx context.Context, c cqrs.Command) error {
		if i >= len(chain) {
			return nil
		}
		current := chain[i]
		i++
		return current(ctx, c, exec)
	}

	return exec(ctx, command)
}

func Register[C cqrs.Command](bus *SimpleCommandBus, handlers ...Middleware) {
	t := typeOf[C]()
	bus.handlers[t] = append(bus.handlers[t], handlers...)
}

func (b *SimpleCommandBus) Use(handlers ...Middleware) {
	b.middlewares = append(b.middlewares, handlers...)
}

func typeOf[C any]() reflect.Type {
	var zero *C
	return reflect.TypeOf(zero).Elem()
}
