package queries

import (
	"app/internal/core/port/cqrs"
	"context"
	"errors"
	"fmt"
	"reflect"
)

type Next func(context.Context, cqrs.Query) (any, error)

type Middleware func(context.Context, cqrs.Query, Next) (any, error)

type SimpleQueryBus struct {
	handlers    map[reflect.Type][]Middleware
	middlewares []Middleware
}

func NewSimpleQueryBus() *SimpleQueryBus {
	return &SimpleQueryBus{handlers: make(map[reflect.Type][]Middleware)}
}

func NewQueryBus(bus *SimpleQueryBus) cqrs.QueryBus {
	return bus
}

func (b *SimpleQueryBus) Execute(ctx context.Context, q cqrs.Query) (any, error) {
	t := reflect.TypeOf(q)
	chain, ok := b.handlers[t]
	if !ok {
		return nil, errors.New(fmt.Sprintf("no handlers found for query %T", q))
	}

	// prepend global middlewares
	chain = append(b.middlewares, chain...)

	// build the chain runner
	i := 0
	var exec Next
	exec = func(ctx context.Context, c cqrs.Query) (any, error) {
		if i >= len(chain) {
			return nil, nil
		}
		current := chain[i]
		i++
		return current(ctx, c, exec)
	}

	return exec(ctx, q)
}

func (b *SimpleQueryBus) Use(handlers ...Middleware) {
	b.middlewares = append(b.middlewares, handlers...)
}

func Register[Q cqrs.Query](bus *SimpleQueryBus, handlers ...Middleware) {
	t := typeOf[Q]()
	bus.handlers[t] = append(bus.handlers[t], handlers...)
}

func Execute[R any](ctx context.Context, bus cqrs.QueryBus, q cqrs.Query) (R, error) {
	r, err := bus.Execute(ctx, q)

	if err != nil {
		var zero R
		return zero, err
	}

	return r.(R), nil
}

func typeOf[Q any]() reflect.Type {
	var zero *Q
	return reflect.TypeOf(zero).Elem()
}
