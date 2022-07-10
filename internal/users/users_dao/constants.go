package users_dao

import "github.com/pkg/errors"

var (
	ErrQueryingDatabase       = errors.New("error querying database")
	ErrParsingRowFromDatabase = errors.New("error parsing row from database")
	ErrHashingPassword        = errors.New("error hashing password")
	ErrSavingToDatabase       = errors.New("error saving to database")
	ErrUpdatingToDatabase     = errors.New("error updating to database")
	ErrDeletingFromDatabase   = errors.New("error deleting from database")
)
