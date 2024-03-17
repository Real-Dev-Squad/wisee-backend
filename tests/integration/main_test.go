package integration_tests

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/Real-Dev-Squad/wisee-backend/src/config"
	"github.com/Real-Dev-Squad/wisee-backend/src/utils"
	"github.com/uptrace/bun"
)

var db *bun.DB

func TestMain(m *testing.M) {
	dsn := config.TestDbUrl
	db = utils.SetupDBConnection(dsn)
	defer TeardownDb(dsn)

	migration_cmd := exec.Command("migrate", "-path", "../../database/migrations", "-database", dsn, "up") // run migrations
	// Execute the migrations
	_, err := migration_cmd.Output()
	if err != nil {
		fmt.Println("Error executing migration up command:", err)
		return
	}
	defer db.Close()

	// setup fixtures
	if err := SetupFixtures(db); err != nil {
		fmt.Println("Error setting up fixtures:", err)
		return
	}

	code := m.Run()
	os.Exit(code)
}
