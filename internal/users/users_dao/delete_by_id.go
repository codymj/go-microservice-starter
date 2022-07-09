package users_dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func deleteQuery() string {
	return `
	delete from Users
	where id = $1
    `
}

// DeleteById an existing User from the database by id
func (r *repository) DeleteById(ctx context.Context, id uuid.UUID) error {
	// execute query
	query := deleteQuery()
	_, err := r.db.DB.ExecContext(ctx, query, id.String())
	if err != nil {
		return errors.Wrap(err, ErrDeletingFromDatabase.Error())
	}

	return nil
}
