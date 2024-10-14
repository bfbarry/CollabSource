package routes

import (
	"fmt"
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
	endpoints = append(endpoints, Route{Path: fmt.Sprintf("%s/project_request/{id}", BASE_URL), Handler: self.projectRequests})
	self.routes = endpoints
}

func (self *ProjectRouter) project(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	UUID := r.Header.Get("UUID")
	userId := r.Header.Get("userId")
	// TODO: return
	// var err *errors.Error
	// var op errors.Op = "routes.project"

	switch r.Method {
	case http.MethodGet:
		fmt.Println("GET /project")
		self.controller.GetProjectByID(w, id, UUID, userId)
	case http.MethodPost:
		fmt.Println("POST /project")
		self.controller.CreateProject(w, r, UUID)
	case http.MethodPatch:
		fmt.Println("PATCH /project")
		self.controller.UpdateProject(w, id, r, UUID)
	case http.MethodDelete:
		fmt.Println("DELETE /project")
		self.controller.DeleteProject(w, id, UUID)
	default:
		responseEntity.SendRequest(w, http.StatusMethodNotAllowed, []byte("Method Not Allowed"))
	}
}

func (self *ProjectRouter) projects(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		fmt.Println("GET /projects")
		self.controller.GetProjects(w, r)
	default:
		responseEntity.SendRequest(w, http.StatusMethodNotAllowed, []byte("Method Not Allowed"))
	}
}

func (self *ProjectRouter) projectRequests(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	UUID := r.Header.Get("UUID")

	switch r.Method {
	case http.MethodPost:
		fmt.Println("POST /project_request")
		self.controller.SendProjectRequest(w, r, UUID)
	case http.MethodPatch:
		fmt.Println("Patch /project_request")
		self.controller.RespondToProjectRequest(w, r, id, UUID)
		// case http.MethodPatch:
		// 	fmt.Println("GET /project_request")
		// 	self.controller.GetProjectRequests(w, r, id, UUID)
	}
}
