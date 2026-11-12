package testhelpers

import (
	"encoding/json"
	"io"
	"testing"
)

func GetJSONFromResponse(t *testing.T, body io.Reader) (jsonData []string) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&jsonData)

	if err != nil {
		t.Fatalf("\nUnable to parse response from server %q \ninto slice of JSON, \n'%v'", body, err)
	}
	return
}
