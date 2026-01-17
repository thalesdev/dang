package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/thalesdev/dang/internal/application/queries"
	"github.com/thalesdev/dang/internal/infrastructure/persistence/sqlite"
	"github.com/thalesdev/dang/internal/interface/rest"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(func() sqlite.Dsn { return "file:/Users/thales/workspace/go/dang/test.sqlite" }),
		queries.Module(),
		sqlite.Module(),
		rest.Module(),
	)

	app.Run()
}
