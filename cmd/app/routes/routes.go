package routes

import (
	"fmt"
	"go-microservice-starter/internal/hello"
	"go-microservice-starter/internal/validate"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	_contentType = "Content-Type"
	_jsonHeader  = "application/json"
	_apiVersion  = "/v1"

	_helloPath = _apiVersion + "/hello"
)

// Services here are initialized in /cmd/app/config/config.go for router access
type Services struct {
	ValidatorService validate.Service
	HelloService     hello.Service
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
	if services.HelloService == nil {
		return handler{}, fmt.Errorf("no hello service provided")
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
	r.Router.HandleFunc(_helloPath, h.postHello).Methods(http.MethodPost)

	return nil
}
