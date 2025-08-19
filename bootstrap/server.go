package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// runServer Starts a server by registering an Fx lifecycle hook
func runServer(lc fx.Lifecycle, router *gin.Engine, zap *zap.SugaredLogger) {
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

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return zap.Sync()
		},
	})
}

var RunServer = fx.Module("run server", fx.Invoke(runServer))
