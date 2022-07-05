package user

import (
	"context"
	"github.com/pkg/errors"
	"go-microservice-starter/internal/database"
)

var (
	ErrQueryingDatabase     = errors.New("error querying database")
	ErrParsingRow           = errors.New("error parsing row returned from database")
	ErrHashingPassword      = errors.New("error hashing password")
	ErrSavingToDatabase     = errors.New("error saving to database")
	ErrUpdatingToDatabase   = errors.New("error updating to database")
	ErrDeletingFromDatabase = errors.New("error deleting from database")
)

// service dependencies to inject
type service struct {
	DB *database.Connection
}

// Service contract
type Service interface {
	List(ctx context.Context) ([]User, error)
	GetById(ctx context.Context, id int64) (User, error)
	GetByUsernamePassword(ctx context.Context, un, pass string) (User, error)
	Save(ctx context.Context, user User) (User, error)
	Update(ctx context.Context, user User) (User, error)
	Delete(ctx context.Context, id int64) error
}

// New returns an initialized instance
func New(db *database.Connection) Service {
	return &service{
		DB: db,
	}
}
