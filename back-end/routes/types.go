package routes

import (
	"net/http"
)

const BASE_URL = "/api/v1"

type IRoutes interface {
	GetRoutes() []Route
}

type Route struct {
	Path string
	Handler func(http.ResponseWriter, *http.Request)
}

type Router struct {
	routes     []Route
}

func (self *Router) GetRoutes() []Route {
	return self.routes
}