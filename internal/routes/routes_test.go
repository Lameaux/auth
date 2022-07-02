package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Lameaux/auth/internal/config"
)

func TestGin(t *testing.T) {
	app := config.NewTestApp()

	r := Gin(app)

	tests := []struct {
		method     string
		url        string
		statusCode int
	}{
		{"GET", "/dummy", http.StatusNotFound},

		{"GET", "/", http.StatusOK},
		{"GET", "/health", http.StatusOK},

		{"POST", "/v1/users", http.StatusNotFound},
	}

	for _, tt := range tests {
		request := httptest.NewRequest(tt.method, tt.url, nil)

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		if status := recorder.Code; status != tt.statusCode {
			t.Errorf("Handler returned wrong status code for %s %s. Expected: %d. Got: %d.", tt.method, tt.url, tt.statusCode, status)
		}
	}
}
