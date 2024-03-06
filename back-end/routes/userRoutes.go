package routes

import (
	"net/http"
	"github.com/bfbarry/CollabSource/back-end/server"
	"github.com/bfbarry/CollabSource/back-end/controllers"
)

type UserRoutes struct {
	routes []server.Endpoint
	userHandler controllers.UserHandler
}

func (userRoutes UserRoutes) GetRoutes() []server.Endpoint {
	return userRoutes.routes
}

func BuildUserRoutes() UserRoutes {
	userRoutes := UserRoutes{}
	userRoutes.userHandler = controllers.BuildUserHandler()
	userRoutes.routes = initiateUserRoutes(userRoutes)

	return userRoutes
}

func initiateUserRoutes(userRoutes UserRoutes) []server.Endpoint{
	endpoints := []server.Endpoint{}
	endpoints = append(endpoints, server.Endpoint{Path:"/users", Handler:userRoutes.users})
	endpoints = append(endpoints, server.Endpoint{Path:"/user/{id}", Handler:userRoutes.user})

	return endpoints
}

func (userRoutes UserRoutes) users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// userRoutes.userHandler.GetAllUsers()
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (userRoutes UserRoutes) user(w http.ResponseWriter, r *http.Request) {
	//bad url parm return 400
	switch r.Method {
	case http.MethodGet:
		userRoutes.userHandler.GetUser(" ID HERE ")
	// case http.MethodPost:
	// 	userRoutes.userHandler.CreateUser()
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}