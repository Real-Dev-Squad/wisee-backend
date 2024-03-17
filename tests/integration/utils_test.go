package integration_tests

import (
	"context"
	"encoding/json"

	"github.com/Real-Dev-Squad/wisee-backend/src/dtos"
	"github.com/Real-Dev-Squad/wisee-backend/src/models"
	"github.com/Real-Dev-Squad/wisee-backend/src/utils/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/uptrace/bun"
)

var user *models.User
var form *models.Form
var formMetaData *models.FormMetaData

type TestResponseDto struct {
	Message string              `json:"message"`
	Data    json.RawMessage     `json:"data"`
	Error   *dtos.ErrorResponse `json:"error"`
}

func SetupFixtures(db *bun.DB) error {
	var ctx = context.Background()
	logger.Info("setting up fixtures")
	defer logger.Info("fixtures setup complete")

	userFixture := &models.User{Username: "test_user", Email: "test_user@admin.com", Password: "password"}
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

func TeardownDb(migrate *migrate.Migrate) {
	logger.Info("Running migration down")
	defer logger.Info("Migration down complete")

	if err := migrate.Down(); err != nil {
		logger.Fatal("failed to run migration down: ", err)
	}
}
