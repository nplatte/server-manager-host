package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestManagerServer(t *testing.T) {
	server := httptest.NewTestServer(t, NewManagerHostServer())
	server.Start()

	client := server.Client()

	testCases := []struct {
		name string
		url  string
	}{
		{"Home Page", server.URL + "/home"},
		{"View All Devices", server.URL + "/view/device/all"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			response, _ := client.Get(tc.url)
			want := http.StatusOK
			got := response.StatusCode
			if want != got {
				t.Errorf("\nWant: %v\nGot: %v", want, got)
			}
		})
	}
}
