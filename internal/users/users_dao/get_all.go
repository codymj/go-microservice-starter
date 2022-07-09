package users_dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
)

func listQuery() string {
	return `
	select
		id,
		username,
		email,
		is_verified,
		created_on,
		updated_on
	from
		users
    `
}

// GetAll returns all rows of User from database
func (r *repository) GetAll(ctx context.Context) ([]*User, error) {
	// execute query
	query := listQuery()
	rows, err := r.db.DB.QueryContext(ctx, query)
	if err != nil {
		log.Err(errors.Wrap(err, ErrQueryingDatabase.Error()))
		return nil, errors.Wrap(err, ErrQueryingDatabase.Error())
	}
	defer Close(&err, io.Closer(rows))

	// parse result
	users := make([]*User, 0)
	for rows.Next() {
		var id uuid.UUID
		var username string
		var email string
		var isVerified bool
		var createdOn int64
		var updatedOn int64

		err = rows.Scan(
			&id, &username, &email, &isVerified, &createdOn, &updatedOn,
		)
		if err != nil {
			log.Err(errors.Wrap(err, ErrParsingRowFromDatabase.Error()))
			return nil, errors.Wrap(err, ErrParsingRowFromDatabase.Error())
		}

		user := User{
			Id:        id,
			Username:  username,
			Email:     email,
			CreatedOn: createdOn,
			UpdatedOn: updatedOn,
		}
		users = append(users, &user)
	}

	return users, nil
}
