package server

import (
	"log"
	"fmt"
	"net/http"
	"errors"
	"github.com/bfbarry/CollabSource/back-end/routes"
)


type Server struct {
	mux *http.ServeMux
}

func CreateNewServer() *Server{
	server := Server{}
	server.mux = http.NewServeMux()

	return &server
}

func (s *Server) StartServer() {
	portNum := 8080
	// defer connections.CloseDB(s.mongoClient)() // TODO: verify pattern
	log.Printf("listening on :%d\n", portNum)
	err := http.ListenAndServe(fmt.Sprintf(":%d", portNum), s.mux)
	// log.Println("closed")
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("http.ErrServerClosed: server shut down \n")
	} else if err != nil {
		log.Printf("error starting server: %s\n", err)
		panic(err)
	}
}

func (s *Server) RegisterRoutes(routes routes.IRoutes) {
	for _,r := range routes.GetRoutes() {
		s.mux.HandleFunc(r.Path, r.Handler)
	}
	
}
