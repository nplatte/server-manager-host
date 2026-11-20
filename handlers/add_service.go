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
}
