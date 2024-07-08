package routes

import (
	"net/http"
	"log"
	"github.com/bfbarry/CollabSource/back-end/controllers"
	"github.com/bfbarry/CollabSource/back-end/responseEntity"
	"fmt"
)

type UserRouter struct {
	Router
	controller *controllers.UserController
}

var defaultUserRouter *UserRouter

func GetDefaultUserRouter() *UserRouter {
	return defaultUserRouter
}

func init() {
	defaultUserRouter = &UserRouter{}
	defaultUserRouter.controller = controllers.GetUserController()
	defaultUserRouter.initiateUserRoutes()
}

func (self *UserRouter) initiateUserRoutes() {
	endpoints := []Route{}
	endpoints = append(endpoints, Route{Path: fmt.Sprintf("%s/user/{id}", BASE_URL), Handler: self.user})
	endpoints = append(endpoints, Route{Path: fmt.Sprintf("%s/users", BASE_URL), Handler: self.users})
	self.routes = endpoints
}

func (self *UserRouter) user(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	switch r.Method {
	case http.MethodGet:
		log.Println("GET /user")
		self.controller.GetUserByID(w, id)
	case http.MethodPost:
		log.Println("POST /user")
		self.controller.CreateUser(w, r)
	case http.MethodPatch:
		log.Println("PATCH /user")
		self.controller.UpdateUser(w, id, r)
	case http.MethodDelete:
		log.Println("DELETE /user")
		self.controller.DeleteUser(w, id)
	default:
		responseEntity.SendRequest(w, http.StatusMethodNotAllowed, []byte("Method Not Allowed"))
	}
}

func (self *UserRouter) users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		log.Println("GET /users")
		self.controller.GetUsers(w, r)
	case http.MethodPost:
		log.Println("POST /users")
		self.controller.GetUsersByQuery(w, r)
		//TODO Methods and structs for getting data by query
		// implementation: pass in multiple user IDs
		break
	default:
		responseEntity.SendRequest(w, http.StatusMethodNotAllowed, []byte("Method not allowed"))
	}
}


