package integration_tests

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"testing"

	"github.com/Real-Dev-Squad/wisee-backend/src/models"
	"github.com/Real-Dev-Squad/wisee-backend/src/routes"
	"github.com/Real-Dev-Squad/wisee-backend/src/utils"
	"github.com/uptrace/bun"
)

var db *bun.DB
var user *models.User
var form *models.Form
var formMetaData *models.FormMetaData

func TeardownDb(dsn string) {
	migration_down_cmd := exec.Command("migrate", "-path", "../../database/migrations", "-database", dsn, "down", "-all") // down the
	// Execute the migrations
	_, err := migration_down_cmd.Output()
	if err != nil {
		log.Fatal("Error executing migration down command:", err)
		return
	}
}

func TestMain(m *testing.M) {
	utils.LoadEnv("../../.env")
	dsn := os.Getenv("TEST_DB_URL")
	db = utils.SetupDBConnection(dsn)
	TeardownDb(dsn)

	migration_cmd := exec.Command("migrate", "-path", "../../database/migrations", "-database", dsn, "up") // run migrations
	// Execute the migrations
	_, err := migration_cmd.Output()
	if err != nil {
		log.Fatal("Error executing migration up command:", err)
		return
	}
	defer db.Close()

	// setup fixtures
	err = SetupFixtures(db)
	if err != nil {
		log.Fatal("Error setting up fixtures:", err)
		return
	}

	code := m.Run()
	os.Exit(code)
}

func SetupFixtures(db *bun.DB) error {
	var ctx = context.Background()

	userFixture := &models.User{Username: "test_user", Email: "test_user@admin.com", Password: "password"}
	fmt.Println(userFixture.Id, "here123")
	if _, err := db.NewInsert().Model(userFixture).Exec(ctx); err != nil {
		return err
	}

	formFixture := &models.Form{OwnerId: userFixture.Id, Status: models.DRAFT, Content: models.FormContent{"blocks": []models.Block{{ID: "1", Type: "text", Content: "Hello World", GroupId: "1", Meta: nil, Order: 1}}}, CreatedById: userFixture.Id}
	if _, err := db.NewInsert().Model(formFixture).Exec(ctx); err != nil {
		return err
	}

	formMetaDataFixture := &models.FormMetaData{
		FormId: formFixture.Id,
	}
	if _, err := db.NewInsert().Model(formMetaDataFixture).Exec(ctx); err != nil {
		return err
	}

	user = userFixture
	form = formFixture
	formMetaData = formMetaDataFixture
	return nil
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
