package routes

import (
	"github.com/bfbarry/CollabSource/back-end/controllers"
)

type RouteEnv struct {
	// this struct allows us to compose route handlers with controllers.Env, which
	// contains database connection, and controller methods
	controllersEnv *controllers.Env
}