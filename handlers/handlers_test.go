package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
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
		form := url.Values{}
		form.Set("name", "Linux PC")
		form.Set("ip", "IP ADDR")
		request := httptest.NewRequest(http.MethodPost, "/create/device", strings.NewReader(form.Encode()))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		testhelpers.AssertEqual(t, 1, len(store.GetDevices()))
	})
}
