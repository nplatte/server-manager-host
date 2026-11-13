package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	t.Run("Test HTML written to response", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		HandleMain(response, request)

		if !strings.Contains(response.Body.String(), "<title>") {
			t.Errorf("Title not in response")
		}

	})
}
