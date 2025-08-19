package middleware

import (
	port "app/internal/core/port/cqrs"
	"app/internal/infrastructure/cqrs/queries"
	"context"
)

func ExecuteQuery[Q port.Query, R any](handler port.QueryHandler[Q, R]) queries.Middleware {
	return func(ctx context.Context, query port.Query, next queries.Next) (any, error) {
		return handler.Execute(ctx, query.(Q))
	}
}
