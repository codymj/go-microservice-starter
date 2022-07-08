package users_dao

import (
	"context"
	"github.com/pkg/errors"
	"go-microservice-starter/internal/database"
)

var (
	_errQueryingDatabase       = errors.New("error querying database")
	_errParsingRowFromDatabase = errors.New("error parsing row from database")
	_errHashingPassword        = errors.New("error hashing password")
	_errSavingToDatabase       = errors.New("error saving to database")
	_errUpdatingToDatabase     = errors.New("error updating to database")
	_errDeletingFromDatabase   = errors.New("error deleting from database")
)

// repository dependencies to inject
type repository struct {
	DB *database.Connection
}

// Repository contract
type Repository interface {
	GetAll(ctx context.Context) ([]*User, error)
	GetById(ctx context.Context, id int64) (*User, error)
	GetByUsernamePassword(ctx context.Context, un, pass string) (*User, error)
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
