package middleware

import (
	"app/internal/core/component/user/application/queries"
	"app/internal/core/component/user/application/queries/port"
	. "app/internal/infrastructure/queries"

	"go.opentelemetry.io/otel/trace"
)

func InitQueryBus(
	bus *SimpleQueryBus,
	uq port.UserQueries,
	tracer trace.Tracer,
) {
	bus.Use(Tracing(tracer))

	Register[queries.FindUserByIDQuery](
		bus,
		ExecuteQuery[queries.FindUserByIDQuery, *port.UserDTO](queries.NewFindUserByIDHandler(uq)),
	)
}
