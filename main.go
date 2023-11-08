package main

import (
	"context"
	"net/http"
	"rentals/internal/app"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			app.NewHTTPServer,
			context.Background,
			app.NewConfig,
			app.NewLogger,
			app.NewPostgres,
			app.NewGinEngine,
			app.NewRepository,
			app.NewHandlers,
		),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
