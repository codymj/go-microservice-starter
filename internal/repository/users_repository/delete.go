package users_repository

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
func (r *repository) Delete(ctx context.Context, id int64) error {
	// execute query
	query := deleteQuery()
	_, err := r.DB.DB.ExecContext(ctx, query, id)
	if err != nil {
		log.Err(errors.Wrap(err, _errDeletingFromDatabase.Error()))
		return errors.Wrap(err, _errDeletingFromDatabase.Error())
	}

	return nil
}
