package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bfbarry/CollabSource/back-end/controllers"
	"github.com/bfbarry/CollabSource/back-end/responseEntity"
	"github.com/bfbarry/CollabSource/back-end/server"
)

const BASE_URL = "/api/v1"

type ProjectRoutes struct {
	routes            []server.Endpoint
	projectController *controllers.ProjectController
}

var defaultProjectRoutes *ProjectRoutes

func GetDefaultProjectRoutes() *ProjectRoutes {
	return defaultProjectRoutes
}

func init() {
	defaultProjectRoutes = &ProjectRoutes{}
	defaultProjectRoutes.projectController = controllers.GetProjectController()
	defaultProjectRoutes.routes = initiateProjectRoutes(defaultProjectRoutes)
}

func (self *ProjectRoutes) GetRoutes() []server.Endpoint {
	return self.routes
}

func initiateProjectRoutes(self *ProjectRoutes) []server.Endpoint {
	endpoints := []server.Endpoint{}
	endpoints = append(endpoints, server.Endpoint{Path: fmt.Sprintf("%s/project/{id}", BASE_URL), Handler: self.project})
	endpoints = append(endpoints, server.Endpoint{Path: fmt.Sprintf("%s/projects", BASE_URL), Handler: self.projects})

	return endpoints
}

func (self *ProjectRoutes) project(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id") // must pass an id
	// TODO: return
	// var err *errors.Error
	// var op errors.Op = "routes.project"

	switch r.Method {
	case http.MethodGet:
		log.Println("GET /project")
		self.projectController.GetProjectByID(w, id)
	case http.MethodPost:
		log.Println("POST /project")
		self.projectController.CreateProject(w, r)
	case http.MethodPatch:
		log.Println("PATCH /project")
		self.projectController.UpdateProject(w, id, r)
	case http.MethodDelete:
		log.Println("DELETE /project")
		self.projectController.DeleteProject(w, id)
	default:
		responseEntity.ResponseEntity(w, http.StatusMethodNotAllowed, []byte("Method Not Allowed"))
	}
}

func (self *ProjectRoutes) projects(w http.ResponseWriter, r *http.Request) {
	// TODO: to be used by elasticsearch
	// var res *controllers.Resp
	// var err *errors.Error

	// switch r.Method {
	// case http.MethodGet:
	// 	log.Println("GET /projects")
	// 	pageNum, intErr := queryParamToInt64(r, "pageNum")
	// 	if intErr != nil {
	// 		log.Printf("Error in queryParamToInt64 %s", err)
	// 		http.Error(w, "query param must be int64", http.StatusBadRequest)
	// 	}
	// 	pageSize, intErr := queryParamToInt64(r, "pageSize")
	// 	if intErr != nil {
	// 		log.Printf("Error in queryParamToInt64 %s", err)
	// 		http.Error(w, "query param must be int64", http.StatusBadRequest)
	// 	}
	// 	res, err = self.projectController.GetProjectsByFilter(&r.Body, pageNum, pageSize)
	// 	if err != nil {
	// 		writeJsonError(w, err)
	// 	}
	// }
	// _ = res
	// writeJsonSuccess(w, res)
}
