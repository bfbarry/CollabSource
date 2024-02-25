package server

import (
	"log"
	"fmt"
	"net/http"
	"errors"
	// "os"
	// "github.com/bfbarry/CollabSource/back-end/connections"
	// "go.mongodb.org/mongo-driver/mongo"

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
	// Env *connections.Env
	// mongoClient *mongo.Client
}

func CreateNewServer() *Server{
	server := Server{}
	server.mux = http.NewServeMux()
	
	// client, db := connections.InitDB()
	// env := &connections.Env{DB: db}
	// server.Env = env
	// server.mongoClient = client
	return &server
}

func (s *Server) StartServer() {
	portNum := 8080
	// defer connections.CloseDB(s.mongoClient)() // TODO: verify pattern
	// log.Printf("listening on :%d\n", portNum)
	err := http.ListenAndServe(fmt.Sprintf(":%d", portNum), s.mux)
	// log.Println("closed")
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("http.ErrServerClosed: server shut down \n")
	} else if err != nil {
		log.Printf("error starting server: %s\n", err)
		panic(err)
	}
}

func (s *Server) RegisterRoutes(routes Routes) {
	for _,r := range routes.GetRoutes() {
		s.mux.HandleFunc(r.Path, r.Handler)
	}
	
}
