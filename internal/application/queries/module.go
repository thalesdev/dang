package queries

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("queries", fx.Provide(
		NewFindAllContentHandler,
		NewFindByIdContentHandler,
	))
}
