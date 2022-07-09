package users_dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go-microservice-starter/internal/database"
	"go-microservice-starter/internal/password"
)

var (
	ErrQueryingDatabase       = errors.New("error querying database")
	ErrParsingRowFromDatabase = errors.New("error parsing row from database")
	ErrHashingPassword        = errors.New("error hashing password")
	ErrSavingToDatabase       = errors.New("error saving to database")
	ErrUpdatingToDatabase     = errors.New("error updating to database")
	ErrDeletingFromDatabase   = errors.New("error deleting from database")
)

// repository dependencies to inject
type repository struct {
	db *database.Connection
	ps password.Service
}

// Repository contract
type Repository interface {
	DeleteById(ctx context.Context, id uuid.UUID) error
	GetById(ctx context.Context, id uuid.UUID) (*User, error)
	GetByParams(ctx context.Context, params map[string]string) ([]*User, error)
	Save(ctx context.Context, user *User) (*User, error)
	UpdateById(ctx context.Context, user *User) (*User, error)
}

// New returns an initialized instance
func New(db *database.Connection, ps password.Service) Repository {
	return &repository{
		db: db,
		ps: ps,
	}
}
