package users

import (
	"context"
	"github.com/google/uuid"
	"go-microservice-starter/internal/users/users_dao"
)

// service dependencies to inject
type service struct {
	ur users_dao.Repository
}

// Service contract
type Service interface {
	DeleteById(ctx context.Context, id uuid.UUID) error
	GetById(ctx context.Context, id uuid.UUID) (users_dao.User, error)
	GetByParams(ctx context.Context, params map[string]string) ([]users_dao.User, error)
	Save(ctx context.Context, r PostUsersRequest) (users_dao.User, error)
	UpdateById(ctx context.Context, id uuid.UUID, r PutUsersIdRequest) (users_dao.User, error)
}

// New returns an initialized instance
func New(ur users_dao.Repository) Service {
	return &service{
		ur: ur,
	}
}
