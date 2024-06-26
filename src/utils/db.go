package utils

import (
	"database/sql"

	"github.com/Real-Dev-Squad/wisee-backend/src/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func SetupDBConnection(dsn string) (*sql.DB, *bun.DB) {
	maxOpenConnections := config.DbMaxOpenConnections

	pgDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	pgDB.SetMaxOpenConns(maxOpenConnections)

	bunDbInstance := bun.NewDB(pgDB, pgdialect.New())

	bunDbInstance.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	return pgDB, bunDbInstance
}
