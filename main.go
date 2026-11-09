package main

import "net/http"

type ManagerHostServer struct {
	*http.ServeMux
}

func NewManagerHostServer() (server *ManagerHostServer) {
	server = &ManagerHostServer{ServeMux: http.NewServeMux()}
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	return
}

func main() {

}
