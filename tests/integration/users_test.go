package tests

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"testing"

	"github.com/Real-Dev-Squad/wisee-backend/src/routes"
	"github.com/Real-Dev-Squad/wisee-backend/src/utils"
	"github.com/uptrace/bun"
)

var db *bun.DB

func TestMain(m *testing.M) {
	utils.LoadEnv("../../.env")
	dsn := os.Getenv("TEST_DB_URL")
	db = utils.SetupDBConnection(dsn)

	migration_cmd := exec.Command("migrate", "-path", "../../migrations", "-database", dsn, "up") // run migrations
	// Execute the migrations
	_, err := migration_cmd.Output()
	if err != nil {
		log.Fatal("Error executing command:", err)
		return
	}
	defer db.Close()

	code := m.Run()

	migration_down_cmd := exec.Command("migrate", "-path", "../../migrations", "-database", dsn, "down", "-all") // down the migrations
	// Execute the migrations
	_, err = migration_down_cmd.Output()
	if err != nil {
		log.Fatal("Error executing command:", err)
		return
	}
	os.Exit(code)
}

func TestGetUsers(t *testing.T) {
	router := routes.SetupV1Routes(db)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/v1/users", nil)

	router.ServeHTTP(w, req)

	if err != nil {
		t.Fatal(err)
	}

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}
}
