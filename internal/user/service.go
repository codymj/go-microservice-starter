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
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]user_repository.User, error)
	GetById(ctx context.Context, id int64) (user_repository.User, error)
	GetByUsernamePassword(ctx context.Context, un, pass string) (user_repository.User, error)
	Save(ctx context.Context, r PostUserRequest) (user_repository.User, error)
}

// New returns an initialized instance
func New(ur user_repository.Repository) Service {
	return &service{
		ur: ur,
	}
}
