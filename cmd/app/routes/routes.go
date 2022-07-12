package routes

import (
	"github.com/gorilla/mux"
	"go-microservice-starter/cmd/app/routes/users_handler"
	"go-microservice-starter/cmd/app/util"
)

const (
	ApiVersion = "/v1"
)

// Router for routing requests
type Router struct {
	Router *mux.Router
}

// NewRouter returns Router pointer
func NewRouter() *Router {
	return &Router{}
}

// Setup creates routes for the app
func (r *Router) Setup(services util.Services) error {
	// init router
	r.Router = mux.NewRouter()

	// init handlers
	uh := users_handler.New(services)
	uh.InitRoutes(r.Router, ApiVersion)

	return nil
}
