package users

import (
	"context"
	"go-microservice-starter/internal/users/users_dao"
)

// service dependencies to inject
type service struct {
	ur users_dao.Repository
}

// Service contract
type Service interface {
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]*users_dao.User, error)
	GetById(ctx context.Context, id int64) (*users_dao.User, error)
	GetByParams(ctx context.Context, params map[string]string) ([]*users_dao.User, error)
	Save(ctx context.Context, r PostUsersRequest) (*users_dao.User, error)
}

// New returns an initialized instance
func New(ur users_dao.Repository) Service {
	return &service{
		ur: ur,
	}
}
