package rest

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/thalesdev/dang/internal/interface/rest/config"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("rest",
		fx.Provide(
			config.NewServer,
			NewContentController,
		),
		fx.Invoke(startServer),
		fx.Invoke(registerControllers),
	)
}

func startServer(lc fx.Lifecycle, app *fiber.App) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go app.Listen(":3000")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})
}

func registerControllers(contentController *ContentController) {}
