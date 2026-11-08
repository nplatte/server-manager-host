package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestManagerServer(t *testing.T) {
	server := httptest.NewTestServer(t, NewManagerHostServer())
	server.Start()

	client := server.Client()

	t.Run("Test Server Basics", func(t *testing.T) {
		url := server.URL
		response, _ := client.Get(url)
		want := http.StatusOK
		got := response.StatusCode
		if want != got {
			t.Errorf("\nWant: %v\nGot: %v", want, got)
		}

	})
}
