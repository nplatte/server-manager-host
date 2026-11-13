package handlers

import (
	"net/http"

	"github.com/nplatte/server-manager-host/stores"
)

type CreateDeviceHandler struct {
	store stores.DeviceStore
}

func NewCreateDeviceHandler(store stores.DeviceStore) (handler *CreateDeviceHandler) {
	return &CreateDeviceHandler{store: store}
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
	w.WriteHeader(http.StatusOK)
	h.store.AddDevice(deviceName, deviceIP)
}
