package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/nplatte/server-manager-host/stores"
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
		{"One Device", [][]string{
			{"Linux", "the IP"},
		}},
		{"Two Devices", [][]string{
			{"Linux", "the IP"},
			{"Windows", "New IP"},
		}},
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

			jsonData := GetDeviceJSONFromResponse(t, response.Body)

			if len(jsonData) != len(tc.devices) {
				t.Errorf("\nWant: %v\nGot: %v", len(tc.devices), len(jsonData))
			}
		})
	}
}

func GetDeviceJSONFromResponse(t *testing.T, body io.Reader) (jsonData []stores.Device) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&jsonData)

	if err != nil {
		t.Fatalf("\nUnable to parse response from server %q \ninto slice of JSON, \n'%v'", body, err)
	}
	return
}
