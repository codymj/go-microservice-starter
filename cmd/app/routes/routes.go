package routes

import (
	"fmt"
	"go-microservice-starter/internal/users"
	"go-microservice-starter/internal/validate"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	_contentType = "Content-Type"
	_jsonHeader  = "application/json"
	_apiVersion  = "/v1"
)

// Services here are initialized in /cmd/app/config/config.go for router access
type Services struct {
	ValidatorService validate.Service
	UserService      users.Service
}

// handler is a wrapper for route handlers to access Services
type handler struct {
	Services
}

// newHandler returns new handler
func newHandler(services Services) (handler, error) {
	if services.ValidatorService == nil {
		return handler{}, fmt.Errorf("no validator service provided")
	}
	if services.UserService == nil {
		return handler{}, fmt.Errorf("no users service provided")
	}

	return handler{services}, nil
}

// NewRouter returns Router pointer
func NewRouter() *Router {
	return &Router{}
}

// Setup creates routes for the app
func (r *Router) Setup(services Services) error {
	// init handler
	h, err := newHandler(services)
	if err != nil {
		return err
	}

	// init router
	r.Router = mux.NewRouter()
	setupUserRoutes(r.Router, h)

	return nil
}

// setupUserRoutes sets up routes for /user endpoint
func setupUserRoutes(r *mux.Router, h handler) {
	usersPath := _apiVersion + "/users" // /v1/users
	usersIdPath := usersPath + "/{id}"  // /v1/users/{id}

	r.HandleFunc(usersPath, h.getUsers).Methods(http.MethodGet)
	r.HandleFunc(usersIdPath, h.getUsersId).Methods(http.MethodGet)
	r.HandleFunc(usersPath, h.postUsers).Methods(http.MethodPost)
	r.HandleFunc(usersIdPath, h.putUsersId).Methods(http.MethodPut)
}
