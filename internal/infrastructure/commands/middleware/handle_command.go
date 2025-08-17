package middleware

import (
	port "app/internal/core/port/commands"
	"app/internal/infrastructure/commands"
	"context"
)

// HandleCommand is a generic wrapper for running CQRS commands
// Usage: commands.Register[ExampleCmd](bus, HandleCommand[ExampleCmd](%ExampleCmdHandler{}))
func HandleCommand[C port.Command](handler port.CommandHandler[C]) commands.Middleware {
	return func(ctx context.Context, command port.Command, next commands.Next) error {
		return handler.Handle(ctx, command.(C))
	}
}
