package users

import (
	"context"
	"go-microservice-starter/internal/users/users_repository"
)

// service dependencies to inject
type service struct {
	ur users_repository.Repository
}

// Service contract
type Service interface {
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]*users_repository.User, error)
	GetById(ctx context.Context, id int64) (*users_repository.User, error)
	GetByUsernamePassword(ctx context.Context, un, pass string) (*users_repository.User, error)
	Save(ctx context.Context, r PostUsersRequest) (*users_repository.User, error)
}

// New returns an initialized instance
func New(ur users_repository.Repository) Service {
	return &service{
		ur: ur,
	}
}
