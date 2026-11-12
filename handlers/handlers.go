package handlers

import (
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/nplatte/server-manager-host.git/ui"
)

func HandleMain(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	template := template.Must(template.ParseFS(ui.Files, "html/home.html"))
	template.Execute(w, nil)
}

func HandleViewDevicesAll(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode([]string{})
}
