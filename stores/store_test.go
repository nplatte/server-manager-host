package stores

import (
	"testing"

	testhelpers "github.com/nplatte/server-manager-host.git/test_helpers"
)

func TestStoreAdd(t *testing.T) {
	testCases := []struct {
		name  string
		store DeviceStore
	}{
		{"In Memory Add", NewMemoryStore()},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			testhelpers.AssertEqual(t, 0, len(test.store.GetDevices()))

			test.store.AddDevice("New Device", "192.168.1.1")

			testhelpers.AssertEqual(t, 1, len(test.store.GetDevices()))
		})
	}
}
