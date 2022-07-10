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

	r.HandleFunc(usersPath, h.getUsers).Methods(http.MethodGet)
	r.HandleFunc(usersIdPath, h.getUsersId).Methods(http.MethodGet)
	r.HandleFunc(usersPath, h.postUsers).Methods(http.MethodPost)
	r.HandleFunc(usersIdPath, h.putUsersId).Methods(http.MethodPut)
	r.HandleFunc(usersIdPath, h.deleteUsersId).Methods(http.MethodDelete)
}

// New returns an initialized instance
func New(services util.Services) Handler {
	return &handler{
		services: services,
	}
}
