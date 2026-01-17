package sqlite

import "github.com/jmoiron/sqlx"

type Dsn string

func NewConnection(dsn Dsn) (*sqlx.DB, error) {
	return sqlx.Connect("sqlite3", string(dsn))
}
