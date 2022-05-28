package routes

import "github.com/gorilla/mux"

// Router for routing requests
type Router struct {
	Router *mux.Router
}

// GenericResponse is for sending generic messages to client
type GenericResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
