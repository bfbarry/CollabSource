package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bfbarry/CollabSource/back-end/controllers"
	"github.com/bfbarry/CollabSource/back-end/responseEntity"
)

type ProjectRouter struct {
	Router
	controller *controllers.ProjectController
}

var defaultProjectRouter *ProjectRouter

func GetDefaultProjectRouter() *ProjectRouter {
	return defaultProjectRouter
}

func init() {
	defaultProjectRouter = &ProjectRouter{}
	defaultProjectRouter.controller = controllers.GetProjectController()
	defaultProjectRouter.initiateProjectRoutes()
}

func (self *ProjectRouter) initiateProjectRoutes() {
	endpoints := []Route{}
	endpoints = append(endpoints, Route{Path: fmt.Sprintf("%s/project/{id}", BASE_URL), Handler: self.project})
	endpoints = append(endpoints, Route{Path: fmt.Sprintf("%s/projects", BASE_URL), Handler: self.projects})
	self.routes = endpoints
}

func (self *ProjectRouter) project(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id") // must pass an id
	// TODO: return
	// var err *errors.Error
	// var op errors.Op = "routes.project"

	switch r.Method {
	case http.MethodGet:
		log.Println("GET /project")
		self.controller.GetProjectByID(w, id)
	case http.MethodPost:
		log.Println("POST /project")
		self.controller.CreateProject(w, r)
	case http.MethodPatch:
		log.Println("PATCH /project")
		self.controller.UpdateProject(w, id, r)
	case http.MethodDelete:
		log.Println("DELETE /project")
		self.controller.DeleteProject(w, id)
	default:
		responseEntity.SendRequest(w, http.StatusMethodNotAllowed, []byte("Method Not Allowed"))
	}
}

func (self *ProjectRouter) projects(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		log.Println("GET /projects")
		self.controller.GetProject(w, r)
	case http.MethodPost:
		//TODO Methods and structs for getting data by query
		break
	default:
		responseEntity.SendRequest(w, http.StatusMethodNotAllowed, []byte("Method Not Allowed"))
	}
}
