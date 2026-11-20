package handlers

import (
	"net/http"

	"github.com/nplatte/server-manager-host/caller"
	"github.com/nplatte/server-manager-host/stores"
)

/*

The Add Service Handler is the Handler that is responsible for adding a service that already exists on
the target machine to the list of manageable services by the webserver

*/

type AddServiceHandler struct {
	s stores.ServiceStore
	c caller.Caller
}

func NewAddServiceHandler(store stores.ServiceStore, caller caller.Caller) (handler *AddServiceHandler) {
	return &AddServiceHandler{s: store, c: caller}
}

func (h *AddServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GET(w, r)
	case http.MethodPost:
		h.POST(w, r)
	default:
		w.WriteHeader(http.StatusForbidden)
	}
}

func (h *AddServiceHandler) GET(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *AddServiceHandler) POST(w http.ResponseWriter, r *http.Request) {
	serviceName := r.FormValue("name")

	for _, field := range []string{serviceName} {
		if field == "" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
	}

	_, err := h.c.AddService(serviceName)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusSeeOther)
	h.s.AddService(serviceName)

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
