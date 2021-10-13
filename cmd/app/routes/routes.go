package routes

import (
	"encoding/json"
	"fmt"
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

/*
 * route handler methods
 **************************************************************************************************/

// getHealth performs a health check on the service
func getHealth(w http.ResponseWriter, r *http.Request) {
	response := GenericResponse{
		Status:  "ok",
		Message: "service is healthy",
	}
	b, err := json.Marshal(response)
	if err != nil {
		err = fmt.Errorf("failed to marshal health response: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(_contentType, _jsonHeader)
	_, err = w.Write(b)
}

/*
 * structs
 **************************************************************************************************/

// Router for routing requests
type Router struct {
	Router *mux.Router
}

// GenericResponse is for sending generic messages to client
type GenericResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
