package fx

import (
	"app/internal/core/component/user/application/queries"

	"go.uber.org/fx"
)

var userComponent = fx.Module("user", fx.Provide(
	queries.NewFindUserByIDHandler,
))

var components = fx.Module("components",
	userComponent,
)

var Core = fx.Module("core",
	components,
)
