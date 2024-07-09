package main

import (
	// "log"

	"github.com/bfbarry/CollabSource/back-end/routes"
	"github.com/bfbarry/CollabSource/back-end/server"
	"github.com/bfbarry/CollabSource/back-end/log"
)

func main() {
	// log.SetFlags(log.LstdFlags)
	// Initilize a web server
	server := server.CreateNewServer()

	server.RegisterRoutes(routes.GetDefaultProjectRouter())
	server.RegisterRoutes(routes.GetDefaultUserRouter())

	log.InitLogEngine("", "stdout", log.DEBUG)
	// Start server
	server.StartServer()
}
