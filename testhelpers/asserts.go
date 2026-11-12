package testhelpers

import "testing"

func AssertEqual(t *testing.T, want any, got any) {
	t.Helper()
	if want != got {
		t.Errorf("\nWant: %v\nGot: %v", want, got)
	}
}
