package main

import (
	bootstrap "app/internal/infrastructure/bootstrap/fx"
	"app/internal/infrastructure/persistence/ent/blog"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
)

// RunServer Lifecycle hook to start the server
func RunServer(lc fx.Lifecycle, router *gin.Engine) {
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

func main() {
	app := fx.New(
		bootstrap.Providers,
		fx.Invoke(
			blog.MigrateSchema,
			RunServer,
		),
	)

	app.Run()
}
