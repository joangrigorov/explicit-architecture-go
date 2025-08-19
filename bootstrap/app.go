package bootstrap

import (
	. "app/bootstrap/fx"

	"go.uber.org/fx"
)

func NewApp() *fx.App {
	return fx.New(
		Config,
		Infrastructure,
		Subscribers,
		Core,
		Presentation,
		Migrations,
		RunServer,
	)
}
