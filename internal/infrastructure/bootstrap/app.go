package bootstrap

import (
	"app/internal/infrastructure/framework/validation"
	"app/internal/infrastructure/persistence/ent/blog"
	"app/internal/presentation/web/core"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
)

// runServer Starts a server by registering an Fx lifecycle hook
func runServer(lc fx.Lifecycle, router *gin.Engine) {
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				fmt.Println("Starting server on :8080")
				if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stopping server...")
			return server.Shutdown(ctx)
		},
	})
}

func NewApp() *fx.App {
	return fx.New(
		providers,
		fx.Invoke(
			blog.MigrateSchema,
			core.RegisterRoutes,
			validation.RegisterBindings,
			validation.RegisterTranslations,
			runServer,
		),
	)
}
