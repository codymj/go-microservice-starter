package user

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func deleteQuery() string {
	return `
	delete Users
	where id = $1
    `
}

// Delete an existing User from the database
func (s *service) Delete(ctx context.Context, id int64) error {
	// execute query
	query := deleteQuery()
	_, err := s.DB.DB.ExecContext(ctx, query, id)
	if err != nil {
		log.Err(errors.Wrap(err, _errDeletingFromDatabase.Error()))
		return errors.Wrap(err, _errDeletingFromDatabase.Error())
	}

	return nil
}