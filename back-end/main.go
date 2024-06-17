package main

import (
	"log"

	"github.com/bfbarry/CollabSource/back-end/routes"
	"github.com/bfbarry/CollabSource/back-end/server"
)

func main() {
	log.SetFlags(log.LstdFlags)
	// Initilize a web server
	server := server.CreateNewServer()

	server.RegisterRoutes(routes.GetDefaultProjectRouter())
	server.RegisterRoutes(routes.GetDefaultUserRouter())

	// Start server
	server.StartServer()
}
