package handlers

import (
	. "app/internal/core/component/user/application/commands"
	"app/internal/core/port/idp"
	"app/internal/core/port/logging"
	"app/internal/infrastructure/commands"
	"app/internal/infrastructure/commands/middleware"
	"app/internal/infrastructure/events"
	ent "app/internal/infrastructure/persistence/ent/generated/user"
	"app/internal/infrastructure/persistence/ent/user"

	"go.opentelemetry.io/otel/trace"
)

func Register(
	l logging.Logger,
	tracer trace.Tracer,
	bus *commands.SimpleCommandBus,
	userRepository *user.Repository,
	idp idp.IdentityProvider,
	eventBus *events.SimpleEventBus,
	entClient *ent.Client,
) {
	bus.Use(middleware.Logger(l))
	bus.Use(middleware.Tracing(tracer))

	commands.Register[RegisterUserCommand](bus, HandleRegisterUserCommand(userRepository, eventBus, entClient))
	commands.Register[ConfirmUserCommand](bus, HandleConfirmUserCommand(userRepository, idp, entClient))
	commands.Register[CreateIdPUserCommand](bus, HandleCreateIdPUserCommand(userRepository, idp, entClient))
}
