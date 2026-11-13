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

func TestAddDeviceHandlerGET(t *testing.T) {
	handler := NewCreateDeviceHandler(stores.NewMemoryStore())
	t.Run("Test GET returns status 200", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/create/device", nil)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		testhelpers.AssertEqual(t, http.StatusOK, response.Code)
	})
}

func TestAddDeviceHandlerPOST(t *testing.T) {
	testCases := []struct {
		name        string
		formData    []string
		deviceCount int
		statusCode  int
	}{
		{
			"Test Handler takes post and makes new device",
			[]string{"Device Name", "Device IP"},
			1, http.StatusOK,
		},
		{
			"Test handler errors when missing name",
			[]string{"", "Device IP"},
			0, http.StatusUnprocessableEntity,
		},
		{
			"Test handler errors when missing IP",
			[]string{"Device Name", ""},
			0, http.StatusUnprocessableEntity,
		},
	}
	for _, tc := range testCases {
		store := stores.NewMemoryStore()
		handler := NewCreateDeviceHandler(store)
		t.Run(tc.name, func(t *testing.T) {
			form := url.Values{}
			form.Set("name", tc.formData[0])
			form.Set("ip", tc.formData[1])

			request := httptest.NewRequest(http.MethodPost, "/create/device", strings.NewReader(form.Encode()))
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			response := httptest.NewRecorder()

			handler.ServeHTTP(response, request)

			testhelpers.AssertEqual(t, tc.statusCode, response.Code)
			testhelpers.AssertEqual(t, tc.deviceCount, len(store.GetDevices()))
		})
	}
}
