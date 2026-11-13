package handlers

import (
	"net/http"
	"text/template"

	"github.com/nplatte/server-manager-host/ui"
	"github.com/nplatte/server-manager-host/stores"
)

func HandleMain(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	template := template.Must(template.ParseFS(ui.Files, "html/home.html"))
	template.Execute(w, nil)
}

type CreateDeviceHandler struct {
	store stores.DeviceStore
}

func NewCreateDeviceHandler(store stores.DeviceStore) (handler *CreateDeviceHandler) {
	return &CreateDeviceHandler{store:store}
}

func (h *CreateDeviceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){

}
