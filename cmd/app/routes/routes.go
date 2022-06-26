package routes

import (
	"fmt"
	"go-microservice-starter/internal/greeting"
	"go-microservice-starter/internal/validate"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	_contentType = "Content-Type"
	_jsonHeader  = "application/json"
	_apiVersion  = "/v1"

	_greetingPath = _apiVersion + "/greeting"
)

// Services here are initialized in /cmd/app/config/config.go for router access
type Services struct {
	ValidatorService validate.Service
	GreetingService  greeting.Service
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
	if services.GreetingService == nil {
		return handler{}, fmt.Errorf("no greeting service provided")
	}

	return handler{services}, nil
}

// NewRouter returns Router pointer
func NewRouter() *Router {
	return &Router{}
}

// Setup creates routes for the app
func (r *Router) Setup(services Services) error {
	h, err := newHandler(services)
	if err != nil {
		return err
	}

	r.Router = mux.NewRouter()
	r.Router.HandleFunc(_greetingPath, h.postGreeting).Methods(http.MethodPost)

	return nil
}
