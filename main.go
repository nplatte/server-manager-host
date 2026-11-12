package main

import (
	"net/http"

	"github.com/nplatte/server-manager-host.git/handlers"
)

type ManagerHostServer struct {
	*http.ServeMux
}

func NewManagerHostServer() (server *ManagerHostServer) {
	server = &ManagerHostServer{ServeMux: http.NewServeMux()}
	server.HandleFunc("/home", handlers.HandleMain)
	server.HandleFunc("/view/device/all", handlers.HandleViewDevicesAll)
	return
}

func main() {

}
