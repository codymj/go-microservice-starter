package user

import (
	"context"
	"github.com/pkg/errors"
	"go-microservice-starter/internal/database"
)

var (
	_errQueryingDatabase     = errors.New("error querying database")
	_errParsingRow           = errors.New("error parsing row returned from database")
	_errHashingPassword      = errors.New("error hashing password")
	_errSavingToDatabase     = errors.New("error saving to database")
	_errUpdatingToDatabase   = errors.New("error updating to database")
	_errDeletingFromDatabase = errors.New("error deleting from database")
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
