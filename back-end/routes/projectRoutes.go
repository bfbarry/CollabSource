package routes

import (
	"net/http"
	"strings"
	"fmt"
	"github.com/bfbarry/CollabSource/back-end/controllers"
	"github.com/bfbarry/CollabSource/back-end/server"
	// "fmt"
)
const BASE_URL = "/api/v1/project"

type ProjectRoutes struct {
	routes []server.Endpoint
}

func (pr ProjectRoutes) GetRoutes() []server.Endpoint {
	return pr.routes
}

func BuildProjectRoutes(env *controllers.Env) ProjectRoutes {
	projectRoutes := ProjectRoutes{
		routes: initiateProjectRoutes(env),
	}
	return projectRoutes
}

func initiateProjectRoutes(env *controllers.Env) []server.Endpoint{
	endpoints := []server.Endpoint{}
	routeEnv := RouteEnv{controllersEnv: env}
	endpoints = append(endpoints, server.Endpoint{Path: fmt.Sprintf("%s/get_one/", BASE_URL), Handler: routeEnv.getOneProject})
	endpoints = append(endpoints, server.Endpoint{Path: fmt.Sprintf("%s/get_many", BASE_URL), Handler: routeEnv.getManyProjects})

	return endpoints
}

func (re *RouteEnv) getOneProject(w http.ResponseWriter, r *http.Request) {
	// TODO: logs,  error handling
	id := strings.TrimPrefix(r.URL.Path, fmt.Sprintf("%s/get_one/", BASE_URL)) // TODO: find safer way
	jsonResponse := re.controllersEnv.GetProjectByID(id)
	// TODO: use middleware for json
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (re *RouteEnv) getManyProjects(w http.ResponseWriter, r *http.Request) {
	// TODO: to be used by elasticsearch
	// TODO: add more filters, pagination
	categoryFilter := r.URL.Query().Get("category")
	jsonResponse := re.controllersEnv.GetProjectsByFilter(categoryFilter)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (re *RouteEnv) createProject(w http.ResponseWriter, r *http.Request) {

}

func (re *RouteEnv) updateProject(w http.ResponseWriter, r *http.Request) {
	
}

func (re *RouteEnv) deleteProject(w http.ResponseWriter, r *http.Request) {
	
}