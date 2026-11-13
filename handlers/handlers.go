package handlers

import (
	"net/http"
	"text/template"

	"github.com/nplatte/server-manager-host/ui"
)

func HandleMain(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	template := template.Must(template.ParseFS(ui.Files, "html/home.html"))
	template.Execute(w, nil)
}
