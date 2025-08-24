package fx

import (
	"app/internal/core/component/user/application/queries"
	"app/internal/core/component/user/application/services"
	"app/internal/infrastructure/component/user/mailables"

	"go.uber.org/fx"
)

var Core = fx.Module("core",
	fx.Module("components",
		fx.Module("user", fx.Provide(
			queries.NewFindUserByIDHandler,
			services.NewConfirmationService,
			services.NewMailService,
			mailables.NewConfirmationMail,
		)),
	),
)
