package integration_tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Real-Dev-Squad/wisee-backend/src/routes"
)

func TestGetUsers(t *testing.T) {
	router := routes.SetupV1Routes(db)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/wisee/v1/users", nil)

	router.ServeHTTP(w, req)

	if err != nil {
		t.Fatal(err)
	}

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}
}
