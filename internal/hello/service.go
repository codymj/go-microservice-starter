package hello

import (
	"context"
	"go-microservice-starter/internal/database"
)

// service dependencies to inject
type service struct {
	DB *database.Connection
}

// Service contract
type Service interface {
	SayHello(ctx context.Context, r PostRequest) PostResponse
}

// New returns an initialized instance
func New(db *database.Connection) Service {
	return &service{
		DB: db,
	}
}
