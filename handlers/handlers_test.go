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

func TestViewDevicesEndpoint(t *testing.T) {
	// the view devices endpoint takes a GET request and returns a list of device names, IP addresses, and a list of services
	testCases := []struct {
		name    string
		devices [][]string
	}{
		{"Base Case", [][]string{}},
	}
	for _, tc := range testCases {
		store := stores.NewMemoryStore()
		handler := NewAllDevicesHandler(store)
		for _, device := range tc.devices {
			store.AddDevice(device[0], device[1])
		}
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/view/device/all", nil)
			response := httptest.NewRecorder()

			handler.ServeHTTP(response, request)

			jsonData := testhelpers.GetJSONFromResponse(t, response.Body)

			if len(jsonData) != len(tc.devices) {
				t.Errorf("\nWant: %v\nGot: %v", len(tc.devices), len(jsonData))
			}
		})
	}
}
