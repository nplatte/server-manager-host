package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nplatte/server-manager-host/stores"
)

type AllDevicesHandler struct {
	store stores.DeviceStore
}

func NewAllDevicesHandler(store stores.DeviceStore) *AllDevicesHandler {
	return &AllDevicesHandler{store: store}
}

func (h *AllDevicesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(h.store.GetDevices())
}
