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
	deviceStore := stores.NewMockDeviceStore()
	serviceStore := stores.NewMockServiceStore()
	serverCaller := caller.NewMockCaller()
	server = &ManagerHostServer{ServeMux: http.NewServeMux()}
	server.HandleFunc("/home", handlers.HandleMain)
	server.Handle("/create/device", handlers.NewCreateDeviceHandler(deviceStore, serverCaller))
	server.Handle("/add/service", handlers.NewAddServiceHandler(serviceStore, serverCaller))
	return
}

func main() {

}
