package util

import (
	"go-microservice-starter/internal/users"
	"go-microservice-starter/internal/validate"
)

// Services here are initialized in /cmd/app/config/config.go for router access
type Services struct {
	ValidatorService validate.Service
	UserService      users.Service
}

// GenericResponse is for sending generic messages to client
type GenericResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
