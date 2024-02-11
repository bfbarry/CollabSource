package server

import (
	"net/http"
	 "log"
)

type Endpoint struct {
	Path string
	Handler func(http.ResponseWriter,*http.Request)
}

type Routes interface {
	GetRoutes() []Endpoint
}

type Server struct {
	mux *http.ServeMux
}

func CreateNewServer() *Server{
	server := Server{}
	server.mux = http.NewServeMux()
	return &server
}

func (s *Server) StartServer() {
	log.Fatal(http.ListenAndServe(":8080", s.mux))
}

func (s *Server) RegisterRoutes(routes Routes) {
	for _,r := range routes.GetRoutes() {
		s.mux.HandleFunc(r.Path, r.Handler)
	}
	
}
