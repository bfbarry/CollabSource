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

	// Build and register all routes
	// userRoutes := routes.BuildUserRoutes()
	// server.RegisterRoutes(userRoutes)
	server.RegisterRoutes(routes.GetDefaultProjectRoutes())

	// Start server
	server.StartServer()
}
