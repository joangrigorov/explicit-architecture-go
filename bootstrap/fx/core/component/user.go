package component

import (
	"app/internal/core/component/user/application/queries"
	"app/internal/core/component/user/application/services"
	"app/internal/core/component/user/application/subscribers"
	"app/internal/infrastructure/component/user/mailables"

	"go.uber.org/fx"
)

var User = fx.Module("user",
	fx.Module("queries", fx.Provide(
		queries.NewFindUserByIDHandler,
	)),
	fx.Module("services", fx.Provide(
		services.NewConfirmationService,
		services.NewMailService,
	)),
	fx.Module("mailables", fx.Provide(
		mailables.NewPasswordSetupMail,
	)),
	fx.Module("subscribers", fx.Provide(
		subscribers.NewSendSetPasswordMailSubscriber,
		subscribers.NewConfirmUserSubscriber,
	)),
)
