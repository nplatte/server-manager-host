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

func TestAddServiceHandlerGET(t *testing.T) {
	handler := NewAddServiceHandler(stores.NewMockServiceStore(), caller.NewMockCaller())
	t.Run("Test GET returns status 200", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/add/service", nil)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		testhelpers.AssertEqual(t, http.StatusOK, response.Code)
	})
}

func TestAddServiceHandlerPOST(t *testing.T) {
	testCases := []struct {
		name         string
		formData     []string
		serviceCount int
		statusCode   int
		location     string
	}{
		{
			"Test Handler takes post and adds existing service",
			[]string{"Service Name"},
			1, http.StatusSeeOther, "/home",
		},
		{
			"Test handler errors when missing name",
			[]string{""},
			0, http.StatusUnprocessableEntity, "",
		},
	}
	testCaller := caller.NewMockCaller()
	for _, tc := range testCases {
		store := stores.NewMockServiceStore()
		handler := NewAddServiceHandler(store, testCaller)
		t.Run(tc.name, func(t *testing.T) {
			form := url.Values{}
			form.Set("name", tc.formData[0])

			request := httptest.NewRequest(http.MethodPost, "/create/device", strings.NewReader(form.Encode()))
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			response := httptest.NewRecorder()

			handler.ServeHTTP(response, request)

			testhelpers.AssertEqual(t, tc.statusCode, response.Code)
			testhelpers.AssertEqual(t, response.Header().Get("Location"), tc.location)
			testhelpers.AssertEqual(t, tc.serviceCount, len(store.GetServices()))
		})
	}
}

func TestAddServiceHandlerCaller(t *testing.T) {
	testCases := []struct {
		name                 string
		expectedServiceCount int
		expectecStatusCode   int
		location             string
		cString              string
		cError               error
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
		{
			"Test ServiceNotFound",
			0, http.StatusOK, "",
			"", caller.ErrServiceNotFound,
		},
		{
			"Test Service Restricted",
			0, http.StatusOK, "",
			"", caller.ErrServiceRestricted,
		},
	}
	testCaller := caller.NewMockCaller()
	for _, tc := range testCases {
		store := stores.NewMockServiceStore()
		handler := NewAddServiceHandler(store, testCaller)
		testCaller.SetReturnString(tc.cString)
		testCaller.SetReturnError(tc.cError)
		t.Run(tc.name, func(t *testing.T) {
			form := url.Values{}
			form.Set("name", "Service Name")

			request := httptest.NewRequest(http.MethodPost, "/add/service", strings.NewReader(form.Encode()))
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			response := httptest.NewRecorder()

			handler.ServeHTTP(response, request)

			testhelpers.AssertEqual(t, tc.expectecStatusCode, response.Code)
			testhelpers.AssertEqual(t, response.Header().Get("Location"), tc.location)
			testhelpers.AssertEqual(t, tc.expectedServiceCount, len(store.GetServices()))
		})
	}
}
