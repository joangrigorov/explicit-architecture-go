package bootstrap

import (
	d "app/bootstrap/fx"
	p "app/bootstrap/fx/presentation"
	"app/config/web"

	"go.uber.org/fx"
)

func WebApp() *fx.App {
	return fx.New(
		fx.Module("config", fx.Provide(web.NewConfig)),
		d.Http,
		d.Logging,
		p.Web,
		RunServer,
	)
}
