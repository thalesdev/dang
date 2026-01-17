package sqlite

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/thalesdev/dang/internal/domain/repositories"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("sqlite", fx.Provide(
		NewConnection,
		fx.Annotate(NewContentRepository, fx.As(new(repositories.ContentRepository))),
	),
		fx.Invoke(connectionHandler),
	)
}

func connectionHandler(lc fx.Lifecycle, db *sqlx.DB) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("starting sqlite connection")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("stopping sqlite connection")
			return db.Close()
		},
	})
}
