package user_repository

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func getByIdQuery() string {
	return `
	select
		id,
		username,
		email,
		created_on,
		last_login
	from
		users
	where
		id = $1
    `
}

// GetById returns a single row of User by id from database
func (r *repository) GetById(ctx context.Context, id int64) (*User, error) {
	// execute query
	query := getByIdQuery()
	row := r.DB.DB.QueryRowContext(ctx, query, id)

	// parse result
	var username string
	var email string
	var createdOn int64
	var lastLogin int64

	err := row.Scan(
		&id, &username, &email, &createdOn, &lastLogin,
	)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return nil, nil
	} else if err != nil {
		log.Err(errors.Wrap(err, _errParsingRowFromDatabase.Error()))
		return &User{}, errors.Wrap(err, _errParsingRowFromDatabase.Error())
	}

	user := User{
		Id:        id,
		Username:  username,
		Email:     email,
		CreatedOn: createdOn,
		LastLogin: lastLogin,
	}

	return &user, nil
}
