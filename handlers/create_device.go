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

type CreateDeviceHandler struct {
	store  stores.DeviceStore
	caller caller.Caller
}

func NewCreateDeviceHandler(store stores.DeviceStore, caller caller.Caller) (handler *CreateDeviceHandler) {
	return &CreateDeviceHandler{store: store, caller: caller}
}

func (h *CreateDeviceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GET(w, r)
	case http.MethodPost:
		h.POST(w, r)
	default:
		w.WriteHeader(http.StatusForbidden)
	}
}

func (h *CreateDeviceHandler) GET(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *CreateDeviceHandler) POST(w http.ResponseWriter, r *http.Request) {
	deviceName := r.FormValue("name")
	deviceIP := r.FormValue("ip")

	for _, field := range []string{deviceName, deviceIP} {
		if field == "" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
	}

	_, err := h.caller.AddDevice(deviceIP)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusSeeOther)
	h.store.AddDevice(deviceName, deviceIP)

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
