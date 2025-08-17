package middleware

import (
	port "app/internal/core/port/commands"
	"app/internal/core/port/logging"
	"app/internal/infrastructure/commands"
	"context"
	"fmt"
)

func Logger(logger logging.Logger) commands.Middleware {
	return func(ctx context.Context, command port.Command, next commands.Next) error {
		logger.Debug(fmt.Sprintf("Dispatched command %T", command))
		err := next(ctx, command)
		if err != nil {
			logger.Error(fmt.Sprintf("Command %T error %s", command, err.Error()))
			return err
		}

		logger.Debug(fmt.Sprintf("Command %T finished", command))
		return nil
	}
}
