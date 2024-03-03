package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

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

func BuildProjectRoutes(env *connections.Env) ProjectRoutes {
	projectRoutes := ProjectRoutes{
		routes: initiateProjectRoutes(env),
	}
	return projectRoutes
}
// TODO: TODAY move controlers.Env to connections
func initiateProjectRoutes(env *connections.Env) []server.Endpoint{
	endpoints := []server.Endpoint{}
	conEnv := &controllers.ProjectEnv{Coll: env.DB.Collection("projects")}
	routeEnv := RouteEnv{controllersEnv: conEnv}
	endpoints = append(endpoints, server.Endpoint{Path: fmt.Sprintf("%s/get_one/", BASE_URL), Handler: routeEnv.getOneProject})
	endpoints = append(endpoints, server.Endpoint{Path: fmt.Sprintf("%s/get_many", BASE_URL), Handler: routeEnv.getManyProjects})
	endpoints = append(endpoints, server.Endpoint{Path: fmt.Sprintf("%s/create", BASE_URL), Handler: routeEnv.createProject})
	endpoints = append(endpoints, server.Endpoint{Path: fmt.Sprintf("%s/update/", BASE_URL), Handler: routeEnv.updateProject})
	endpoints = append(endpoints, server.Endpoint{Path: fmt.Sprintf("%s/delete/", BASE_URL), Handler: routeEnv.deleteProject})

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
	// TODO: add user metadata
	log.Println("createProject")
	var p controllers.Project
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jsonResponse := re.controllersEnv.CreateProject(p)
	w.Write(jsonResponse)
}

func (re *RouteEnv) updateProject(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, fmt.Sprintf("%s/update/", BASE_URL))
	var p controllers.Project
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jsonResponse := re.controllersEnv.UpdateProject(id, p)
	w.Write(jsonResponse)
}

func (re *RouteEnv) deleteProject(w http.ResponseWriter, r *http.Request) {
	log.Println("deleteProject")
	id := strings.TrimPrefix(r.URL.Path, fmt.Sprintf("%s/delete/", BASE_URL))
	deleteModeStr := r.URL.Query().Get("mode")
	deleteMode := controllers.Str2Enum(deleteModeStr)
	jsonResponse := re.controllersEnv.DeleteProject(deleteMode, id)
	w.Write(jsonResponse)
}