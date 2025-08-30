package component

import (
	"app/internal/core/component/user/application/queries/find_user_by_id"
	"app/internal/core/component/user/application/queries/get_verification_preflight"
	"app/internal/core/component/user/application/services"
	"app/internal/core/component/user/application/subscribers"
	"app/internal/infrastructure/component/user/mailables"

	"go.uber.org/fx"
)

var User = fx.Module("user",
	fx.Module("queries", fx.Provide(
		find_user_by_id.NewHandler,
		get_verification_preflight.NewHandler,
	)),
	fx.Module("services", fx.Provide(
		services.NewVerificationService,
		services.NewMailService,
	)),
	fx.Module("mailables", fx.Provide(
		mailables.NewPasswordSetupMail,
		mailables.NewUserConfirmedMail,
	)),
	fx.Module("subscribers", fx.Provide(
		subscribers.NewSendSetPasswordMailSubscriber,
		subscribers.NewConfirmUserSubscriber,
	)),
)
