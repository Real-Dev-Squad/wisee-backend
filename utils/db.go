package utils

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func SetupDBConnection(dsn string) *bun.DB {
	pgDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	pgDB.SetMaxOpenConns(1)

	db := bun.NewDB(pgDB, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	return db
}
