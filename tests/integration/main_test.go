package integration_tests

import (
	"testing"

	_ "github.com/lib/pq"

	"github.com/Real-Dev-Squad/wisee-backend/src/config"
	"github.com/Real-Dev-Squad/wisee-backend/src/utils"
	"github.com/Real-Dev-Squad/wisee-backend/src/utils/logger"
	"github.com/uptrace/bun"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var db *bun.DB

func TestMain(m *testing.M) {
	dsn := config.TestDbUrl
	db, bunDb := utils.SetupDBConnection(dsn)
	defer bunDb.Close()

	// use the db connection from the `setupDBConnection` function to run migrations
	driver, pgInstanceErr := postgres.WithInstance(db, &postgres.Config{})

	if pgInstanceErr != nil {
		logger.Fatal("pg instance error: ", pgInstanceErr)
	}

	migration, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations", // "file://" + path of the migrations folder (not using "file://" will throw an error)
		"postgres", driver)

	if err != nil {
		logger.Fatal("migrate: database instance error: ", err)
	}

	if err := migration.Up(); err != nil {
		logger.Fatal("migration error: ", err)
	}

	logger.Info("Migrations complete")

	// teardown the database after the tests
	defer TeardownDb(migration)

	// setup fixtures
	if err := SetupFixtures(bunDb); err != nil {
		logger.Fatal("Error setting up fixtures:", err)
	}
}
