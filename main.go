package main

import (
	"net/http"

	"github.com/nplatte/server-manager-host/caller"
	"github.com/nplatte/server-manager-host/handlers"
	"github.com/nplatte/server-manager-host/stores"
)

type ManagerHostServer struct {
	*http.ServeMux
}

func NewManagerHostServer() (server *ManagerHostServer) {
	store := stores.NewMockDeviceStore()
	serverCaller := caller.NewMockCaller()
	server = &ManagerHostServer{ServeMux: http.NewServeMux()}
	server.HandleFunc("/home", handlers.HandleMain)
	server.Handle("/create/device", handlers.NewCreateDeviceHandler(store, serverCaller))
	server.Handle("/add/service", handlers.NewAddServiceHandler(store, serverCaller))
	return
}

func main() {

}
