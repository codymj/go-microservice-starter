package users_handler

import (
	"github.com/gorilla/mux"
	"go-microservice-starter/cmd/app/util"
	"net/http"
)

type handler struct {
	services util.Services
}

type Handler interface {
	InitRoutes(r *mux.Router, apiVersion string)
}

func (h *handler) InitRoutes(r *mux.Router, apiVersion string) {
	usersPath := apiVersion + "/users"
	usersIdPath := usersPath + "/{id}"

	r.HandleFunc(usersIdPath, h.getById).Methods(http.MethodGet)
	r.HandleFunc(usersPath, h.getByParams).Methods(http.MethodGet)
	r.HandleFunc(usersPath, h.post).Methods(http.MethodPost)
	r.HandleFunc(usersIdPath, h.putById).Methods(http.MethodPut)
	r.HandleFunc(usersIdPath, h.deleteById).Methods(http.MethodDelete)
}

// New returns an initialized instance
func New(services util.Services) Handler {
	return &handler{
		services: services,
	}
}
