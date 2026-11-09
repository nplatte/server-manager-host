package main

import (
	"encoding/json"
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

func TestViewDevicesEndpoint(t *testing.T) {
	// the view devices endpoint takes a GET request and returns a list of device names, IP addresses, and a list of services
	testCases := []struct {
		name  string
		count int
	}{
		{"Base Case", 0},
	}
	for _, tc := range testCases {
		for i := range tc.count {
			i = i
		}
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/view/device/all", nil)
			response := httptest.NewRecorder()

			HandleViewDevicesAll(response, request)

			var jsonData []string
			err := json.NewDecoder(response.Body).Decode(&jsonData)

			if err != nil {
				t.Fatalf("\nUnable to parse response from server %q \ninto slice of JSON, \n'%v'", response.Body, err)
			}

			if len(jsonData) != tc.count {
				t.Errorf("\nWant: %v\nGot: %v", tc.count, len(jsonData))
			}
		})
	}
}
