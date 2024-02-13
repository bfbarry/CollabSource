package routes

import (
	"net/http"
	"github.com/bfbarry/CollabSource/back-end/server"
	"github.com/bfbarry/CollabSource/back-end/controllers"
	"fmt"
)

type UserRoutes struct {
	routes []server.Endpoint // TODO make this private
}

func (userRoutes UserRoutes) GetRoutes() []server.Endpoint {
	return userRoutes.routes
}

func BuildUserRoutes(env *controllers.Env) UserRoutes {
	userRoutes := UserRoutes{}
	userRoutes.routes = initiateUserRoutes(env)
	return userRoutes
}

func initiateUserRoutes(env *controllers.Env) []server.Endpoint{
	endpoints := []server.Endpoint{}
	conEnv := &controllers.ProjectEnv{Coll: env.DB.Collection("projects")}
	routeEnv := RouteEnv{controllersEnv: conEnv}
	endpoints = append(endpoints, server.Endpoint{Path:"/user", Handler:routeEnv.getUserByID})

	return endpoints
}

func (re *RouteEnv) getUserByID(w http.ResponseWriter, r *http.Request) {
	// TODO switch statment 
	fmt.Fprintf(w, "Hi there, I love you!")
}

func (re *RouteEnv) user(w http.ResponseWriter, r *http.Request) {
	// TODO switch statment 
	fmt.Fprintf(w, "Hi there, I love you!")
}