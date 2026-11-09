package main

import (
	"fmt"
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
		{"Home Page", fmt.Sprintf("%v/home", server.URL)},
		{"View All Devices", fmt.Sprintf("%v/view/device/none", server.URL)},
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
