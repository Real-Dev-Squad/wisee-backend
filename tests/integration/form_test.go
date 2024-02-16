package integration_tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Real-Dev-Squad/wisee-backend/src/dtos"
	"github.com/Real-Dev-Squad/wisee-backend/src/models"
	"github.com/Real-Dev-Squad/wisee-backend/src/routes"
)

type TestResponseDto struct {
	Message string             `json:"message"`
	Data    json.RawMessage    `json:"data"`
	Error   dtos.ErrorResponse `json:"error"`
}

func TestFormCreate(t *testing.T) {
	router := routes.SetupV1Routes(db)
	// add the DTO
	var requestBody = map[string]interface{}{
		"status":          models.DRAFT,
		"performed_by_id": user.Id,
		"content":         models.FormContent{"blocks": []models.Block{{ID: "1", Type: "text", Content: "Hello World", GroupId: "1", Meta: nil, Order: 1}}},
	}

	// Convert requestBody to JSON
	jsonValue, _ := json.Marshal(requestBody)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/v1/forms", bytes.NewBuffer(jsonValue))

	router.ServeHTTP(w, req)

	if err != nil {
		t.Fatal(err)
	}

	var respBody TestResponseDto
	if err := json.NewDecoder(w.Body).Decode(&respBody); err != nil {
		t.Fatal(err)
	}

	var resData dtos.CreateUpdateGetFormResponseDto
	if err := json.Unmarshal(respBody.Data, &resData); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "form created successfully", respBody.Message)

	var ctx = context.Background()
	var formMetaData models.FormMetaData
	if err := db.NewSelect().Model(&formMetaData).Where("form_id = ?", resData.Id).Scan(ctx); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, resData.Id, formMetaData.FormId)
}

func TestFormGetAll(t *testing.T) {
	router := routes.SetupV1Routes(db)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/v1/forms", nil)

	router.ServeHTTP(w, req)

	if err != nil {
		t.Fatal(err)
	}

	var respBody TestResponseDto
	if err := json.NewDecoder(w.Body).Decode(&respBody); err != nil {
		t.Fatal(err)
	}

	var resData dtos.GetFormsResponseDto
	if err := json.Unmarshal(respBody.Data, &resData); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "forms fetched successfully", respBody.Message)
	assert.LessOrEqual(t, 1, len(resData))
}

func TestFormGetById(t *testing.T) {
	router := routes.SetupV1Routes(db)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/v1/forms/%v", form.Id), nil)

	router.ServeHTTP(w, req)
	if err != nil {
		t.Fatal(err)
	}

	var respBody TestResponseDto
	if err := json.NewDecoder(w.Body).Decode(&respBody); err != nil {
		t.Fatal(err)
	}

	var resData dtos.GetFormDetailResponseDto
	if err := json.Unmarshal(respBody.Data, &resData); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "form fetched successfully", respBody.Message)
	assert.Equal(t, form.Id, resData.Id)
	assert.Equal(t, form.OwnerId, resData.OwnerId)
	assert.Equal(t, string(form.Status), resData.Status)
	assert.Equal(t, form.CreatedById, resData.CreatedById)
	assert.Equal(t, formMetaData.Id, resData.Meta.Id)
	assert.Equal(t, formMetaData.FormId, resData.Meta.FormId)
}

func TestFormUpdate(t *testing.T) {
	assert.Nil(t, form.UpdatedById)

	router := routes.SetupV1Routes(db)

	// add the DTO
	var requestBody = map[string]interface{}{
		"status":          models.DRAFT,
		"performed_by_id": user.Id,
		"content":         models.FormContent{"blocks": []models.Block{{ID: "1", Type: "text", Content: "Hello World", GroupId: "1", Meta: nil, Order: 1}}},
	}

	// Convert requestBody to JSON
	jsonValue, _ := json.Marshal(requestBody)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("PATCH", fmt.Sprintf("/v1/forms/%v", form.Id), bytes.NewBuffer(jsonValue))

	router.ServeHTTP(w, req)
	if err != nil {
		t.Fatal(err)
	}

	var respBody TestResponseDto
	if err := json.NewDecoder(w.Body).Decode(&respBody); err != nil {
		t.Fatal(err)
	}

	var resData dtos.CreateUpdateGetFormResponseDto
	if err := json.Unmarshal(respBody.Data, &resData); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusAccepted, w.Code)
	assert.Equal(t, "form updated successfully", respBody.Message)
	assert.Equal(t, user.Id, *resData.UpdatedById)
}
