package routes

import (
	"net/http"
	"github.com/bfbarry/CollabSource/back-end/server"
	"fmt"
)

type UserRoutes struct {
	Routes []server.Endpoint
}

func (userRoutes UserRoutes) GetRoutes() []server.Endpoint {
	return userRoutes.Routes
}

func BuildUserRoutes() UserRoutes {
	userRoutes := UserRoutes{}
	userRoutes.Routes = initiateRoutes()
	return userRoutes
}

func initiateRoutes() []server.Endpoint{
	endpoints := []server.Endpoint{}
	endpoints = append(endpoints, server.Endpoint{Path:"/user", Handler:getUserByID})

	return endpoints
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	// TODO switch statment 
	fmt.Fprintf(w, "Hi there, I love you!")
}

func user(w http.ResponseWriter, r *http.Request) {
	// TODO switch statment 
	fmt.Fprintf(w, "Hi there, I love you!")
}