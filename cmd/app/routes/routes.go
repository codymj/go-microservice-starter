package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	_contentType = "Content-Type"
	_jsonHeader  = "application/json"

	_api = "/api"
	_v1  = "/v1"
)

// NewRouter returns Router pointer
func NewRouter() *Router {
	return &Router{}
}

// Setup creates routes for the app
func (r *Router) Setup() {
	log.Println("initializing routes...")

	r.Router = mux.NewRouter()
	r.Router.HandleFunc(_api+_v1+"/health", getHealth).Methods(http.MethodGet)
}
