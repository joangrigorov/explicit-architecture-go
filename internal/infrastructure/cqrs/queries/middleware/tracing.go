package middleware

import (
	port "app/internal/core/port/cqrs"
	"app/internal/infrastructure/cqrs/queries"
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func Tracing(tracer trace.Tracer) queries.Middleware {
	return func(ctx context.Context, query port.Query, next queries.Next) (any, error) {
		ctx, span := tracer.Start(ctx, fmt.Sprintf("Query %T", query))
		defer span.End()
		payload, _ := query.LogBody()
		span.AddEvent(
			fmt.Sprintf("%T params", query),
			trace.WithAttributes(attribute.String("params", string(payload))),
		)

		res, err := next(ctx, query)

		if err != nil {
			span.RecordError(err)
		}

		return res, err
	}
}
