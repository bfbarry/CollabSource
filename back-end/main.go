package main

import (
	"github.com/bfbarry/CollabSource/back-end/server"
	"github.com/bfbarry/CollabSource/back-end/routes"
)

func main() {
	
	// Initilize a web server
	server := server.CreateNewServer()

	// Build and register all routes
	userRoutes := routes.BuildUserRoutes()
	server.RegisterRoutes(userRoutes)
	//// projectRoutes := routes.BuildProjectRoutes(server.Env)
	//// server.RegisterRoutes(projectRoutes)

	// Start server
	server.StartServer()
}