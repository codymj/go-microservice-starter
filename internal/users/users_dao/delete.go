package users_dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func deleteQuery() string {
	return `
	delete from Users
	where id = $1
    `
}

// Delete an existing User from the database
func (r *repository) Delete(ctx context.Context, id uuid.UUID) error {
	// execute query
	query := deleteQuery()
	_, err := r.db.DB.ExecContext(ctx, query, id.String())
	if err != nil {
		log.Err(errors.Wrap(err, ErrDeletingFromDatabase.Error()))
		return errors.Wrap(err, ErrDeletingFromDatabase.Error())
	}

	return nil
}
