package routes

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"go-microservice-starter/internal/hello"
	"go-microservice-starter/internal/validate"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	_contentType = "Content-Type"
	_jsonHeader  = "application/json"

	_api = "/api"
	_v1  = "/v1"
)

type Services struct {
	ValidatorService validate.Service
	HelloService     hello.Service
}

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
	log.Info().Msg("initializing routes...")

	h, err := newHandler(services)
	if err != nil {
		return err
	}

	r.Router = mux.NewRouter()
	r.Router.HandleFunc(_api+_v1+"/hello", h.postHello).Methods(http.MethodPost)

	return nil
}
