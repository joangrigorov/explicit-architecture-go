package middleware

import (
	port "app/internal/core/port/commands"
	"app/internal/infrastructure/commands"
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func Tracing(tracer trace.Tracer) commands.Middleware {
	return func(ctx context.Context, command port.Command, next commands.Next) error {
		ctx, span := tracer.Start(ctx, fmt.Sprintf("Command %T", command))
		defer span.End()
		payload, _ := command.LogBody()
		span.AddEvent(
			fmt.Sprintf("%T payload", command),
			trace.WithAttributes(attribute.String("payload", string(payload))),
		)

		if err := next(ctx, command); err != nil {
			span.RecordError(err)
			return err
		}

		return nil
	}
}
