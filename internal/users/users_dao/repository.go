package users_dao

import (
	"context"
	"github.com/pkg/errors"
	"go-microservice-starter/internal/database"
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
	DB *database.Connection
}

// Repository contract
type Repository interface {
	GetAll(ctx context.Context) ([]*User, error)
	GetById(ctx context.Context, id int64) (*User, error)
	GetByParams(ctx context.Context, params map[string]string) ([]*User, error)
	Save(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, id int64) error
}

// New returns an initialized instance
func New(db *database.Connection) Repository {
	return &repository{
		DB: db,
	}
}
