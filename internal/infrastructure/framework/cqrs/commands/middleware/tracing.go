package middleware

import (
	port "app/internal/core/port/cqrs"
	"app/internal/infrastructure/framework/cqrs/commands"
	"context"
	"fmt"
	"reflect"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func Tracing(tracer trace.Tracer) commands.Middleware {
	return func(ctx context.Context, command port.Command, next commands.Next) error {
		ctx, span := tracer.Start(ctx, fmt.Sprintf("Command %T", command))
		defer span.End()

		span.SetAttributes(
			attribute.String("command", id(command)),
		)

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

func id(e interface{}) string {
	t := reflect.TypeOf(e)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t.PkgPath() + "." + t.Name()
}
