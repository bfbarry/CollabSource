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
	defaultUserRouter.routes = initiateUserRoutes(defaultUserRouter)
}

func initiateUserRoutes(self *UserRouter) []Route {
	endpoints := []Route{}
	endpoints = append(endpoints, Route{Path: fmt.Sprintf("%s/user/{id}", BASE_URL), Handler: self.user})
	return endpoints
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
