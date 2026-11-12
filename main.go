package main

import (
	"net/http"

	"github.com/nplatte/server-manager-host/handlers"
	"github.com/nplatte/server-manager-host/stores"
)

type ManagerHostServer struct {
	*http.ServeMux
}

func NewManagerHostServer() (server *ManagerHostServer) {
	store := stores.NewMemoryStore()
	server = &ManagerHostServer{ServeMux: http.NewServeMux()}
	server.HandleFunc("/home", handlers.HandleMain)
	server.Handle("/view/device/all", handlers.NewAllDevicesHandler(store))
	return
}

func main() {

}
