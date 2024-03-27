package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bfbarry/CollabSource/back-end/controllers"
	// "github.com/bfbarry/CollabSource/back-end/repository"
	"github.com/bfbarry/CollabSource/back-end/server"
	"github.com/bfbarry/CollabSource/back-end/errors"
)
const BASE_URL = "/api/v1"

type ProjectRoutes struct {
	routes []server.Endpoint
	handler *controllers.ProjectHandler // TODO rename to controller
}

var DefaultProjectRoutes *ProjectRoutes

func (self *ProjectRoutes) GetRoutes() []server.Endpoint {
	return self.routes
}

func init() {
	DefaultProjectRoutes = &ProjectRoutes{}
	DefaultProjectRoutes.handler = controllers.BuildProjectHandler()
	DefaultProjectRoutes.routes = initiateProjectRoutes(DefaultProjectRoutes)
}

func initiateProjectRoutes(self *ProjectRoutes) []server.Endpoint{
	endpoints := []server.Endpoint{}
	endpoints = append(endpoints, server.Endpoint{Path: fmt.Sprintf("%s/project/{id}", BASE_URL), Handler: self.project})
	endpoints = append(endpoints, server.Endpoint{Path: fmt.Sprintf("%s/projects", BASE_URL), Handler: self.projects})

	return endpoints
}

// TODO: separate methods in functions
func (self *ProjectRoutes) project(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id") // must pass an id
	// TODO: return 
	var res []byte
	var err *errors.Error
	var op errors.Op = "routes.project"

	switch r.Method {
	case http.MethodGet:
		log.Println("GET /project")
		res, err = self.handler.GetProjectByID(id)
	case http.MethodPost:
		log.Println("POST /project")
		res, err = self.handler.CreateProject(&r.Body)
	// case http.MethodPatch:
	// 	log.Println("PATCH /project")
	// 	res, err = self.handler.UpdateProject(id, &r.Body)
	// case http.MethodDelete:
	// 	log.Println("DELETE /project")
	// 	deleteModeStr := r.URL.Query().Get("mode")
	// 	deleteMode := repository.Str2Enum(deleteModeStr)
	// 	res, err = self.handler.DeleteProject(deleteMode, id)
	default:
		writeJsonError(w, errors.E(err, http.StatusMethodNotAllowed, op, "Method Not Allowed"))
	}

	if err != nil {
		writeJsonError(w, err)
	}
	writeJsonSuccess(w, res)
}

func (self *ProjectRoutes) projects(w http.ResponseWriter, r *http.Request) {
	// TODO: to be used by elasticsearch
	var res []byte
	var err error

	switch r.Method {
	case http.MethodGet:
		log.Println("GET /projects")
		pageNum, intErr := queryParamToInt64(r, "pageNum")
		if intErr != nil {
			log.Printf("Error in queryParamToInt64 %s", err)
			http.Error(w, "query param must be int64", http.StatusBadRequest)
		}
		pageSize, intErr := queryParamToInt64(r, "pageSize")
		if intErr != nil {
			log.Printf("Error in queryParamToInt64 %s", err)
			http.Error(w, "query param must be int64", http.StatusBadRequest)
		}
		res, err = self.handler.GetProjectsByFilter(&r.Body, pageNum, pageSize)
		if err != nil {
			log.Printf("Error in GetProjectsByFilter %s", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	} 
	writeJsonSuccess(w, res)
}