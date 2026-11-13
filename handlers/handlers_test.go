package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/nplatte/server-manager-host/stores"
	"github.com/nplatte/server-manager-host/testhelpers"
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


func TestAddDeviceHandler(t *testing.T) {
	store := stores.NewMemoryStore()
	handler := NewCreateDeviceHandler(store)
	t.Run("Test Handler takes post and makes new device", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/create/device", nil)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		testhelpers.AssertEqual(t, 1, len(store.GetDevices()))
	})
}
