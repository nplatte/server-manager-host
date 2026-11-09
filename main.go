package main

import "net/http"

type ManagerHostServer struct {
	*http.ServeMux
}

func NewManagerHostServer() (server *ManagerHostServer) {
	server = &ManagerHostServer{ServeMux: http.NewServeMux()}
	server.HandleFunc("/home", HandleMain)
	return
}

func main() {

}
