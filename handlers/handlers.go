package handlers

import (
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/nplatte/server-manager-host/stores"
	"github.com/nplatte/server-manager-host/ui"
)

func HandleMain(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	template := template.Must(template.ParseFS(ui.Files, "html/home.html"))
	template.Execute(w, nil)
}

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
