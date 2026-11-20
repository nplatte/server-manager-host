package handlers

import (
	"net/http"

	"github.com/nplatte/server-manager-host/caller"
	"github.com/nplatte/server-manager-host/stores"
)

/*

The Create Device Handler is the Handler that is responsible for adding a device to the service listed.
It tells the caller to add the device to it's open connections
It also saves the data to the database

*/

type AddServiceHandler struct {
	s stores.ServiceStore
	c caller.Caller
}

func NewAddServiceHandler(store stores.DeviceStore, caller caller.Caller) (handler *AddServiceHandler) {
	return &AddServiceHandler{s: store, c: caller}
}

func (h *AddServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
