package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/nplatte/server-manager-host/caller"
	"github.com/nplatte/server-manager-host/stores"
	"github.com/nplatte/server-manager-host/testhelpers"
)

func TestAddDeviceHandlerGET(t *testing.T) {
	handler := NewCreateDeviceHandler(stores.NewMockDeviceStore(), caller.NewMockCaller())
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
		location    string
		cString     string
		cError      error
	}{
		{
			"Test Handler takes post and makes new device",
			[]string{"Device Name", "Device IP"},
			1, http.StatusSeeOther, "/home",
			"", nil,
		},
		{
			"Test handler errors when missing name",
			[]string{"", "Device IP"},
			0, http.StatusUnprocessableEntity, "",
			"", nil,
		},
		{
			"Test handler errors when missing IP",
			[]string{"Device Name", ""},
			0, http.StatusUnprocessableEntity, "",
			"", nil,
		},
	}
	testCaller := caller.NewMockCaller()
	for _, tc := range testCases {
		store := stores.NewMockDeviceStore()
		handler := NewCreateDeviceHandler(store, testCaller)
		testCaller.SetReturnString(tc.cString)
		testCaller.SetReturnError(tc.cError)
		t.Run(tc.name, func(t *testing.T) {
			form := url.Values{}
			form.Set("name", tc.formData[0])
			form.Set("ip", tc.formData[1])

			request := httptest.NewRequest(http.MethodPost, "/create/device", strings.NewReader(form.Encode()))
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			response := httptest.NewRecorder()

			handler.ServeHTTP(response, request)

			testhelpers.AssertEqual(t, tc.statusCode, response.Code)
			testhelpers.AssertEqual(t, response.Header().Get("Location"), tc.location)
			testhelpers.AssertEqual(t, tc.deviceCount, len(store.GetDevices()))
		})
	}
}

func TestAddDeviceHandlerCaller(t *testing.T) {
	testCases := []struct {
		name                string
		expectedDeviceCount int
		expectecStatusCode  int
		location            string
		cString             string
		cError              error
	}{
		{
			"Test Caller returns Success",
			1, http.StatusSeeOther, "/home",
			"", nil,
		},
		{
			"Test Caller Can't Reach Device",
			0, http.StatusOK, "",
			"", caller.ErrUnreacable,
		},
		{
			"Test Caller got ConnRefused",
			0, http.StatusOK, "",
			"", caller.ErrRefused,
		},
	}
	testCaller := caller.NewMockCaller()
	for _, tc := range testCases {
		store := stores.NewMockDeviceStore()
		handler := NewCreateDeviceHandler(store, testCaller)
		testCaller.SetReturnString(tc.cString)
		testCaller.SetReturnError(tc.cError)
		t.Run(tc.name, func(t *testing.T) {
			form := url.Values{}
			form.Set("name", "Device Name")
			form.Set("ip", "Device IP")

			request := httptest.NewRequest(http.MethodPost, "/create/device", strings.NewReader(form.Encode()))
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			response := httptest.NewRecorder()

			handler.ServeHTTP(response, request)

			testhelpers.AssertEqual(t, tc.expectecStatusCode, response.Code)
			testhelpers.AssertEqual(t, response.Header().Get("Location"), tc.location)
			testhelpers.AssertEqual(t, tc.expectedDeviceCount, len(store.GetDevices()))
		})
	}
}
