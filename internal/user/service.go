package user

import (
	"context"
	"go-microservice-starter/internal/repository/user_repository"
)

// service dependencies to inject
type service struct {
	ur user_repository.Repository
}

// Service contract
type Service interface {
	GetAll(ctx context.Context) ([]user_repository.User, error)
	Create(ctx context.Context, r PostUserRequest) (user_repository.User, error)
}

// New returns an initialized instance
func New(ur user_repository.Repository) Service {
	return &service{
		ur: ur,
	}
}
