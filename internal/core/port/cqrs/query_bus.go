package cqrs

import "context"

type Query interface {
	LogBody() ([]byte, error)
}

type QueryHandler[Q Query, R any] interface {
	Execute(context.Context, Q) (R, error)
}

type QueryBus interface {
	Execute(context.Context, Query) (any, error)
}
