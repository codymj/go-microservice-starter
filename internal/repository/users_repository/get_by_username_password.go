package users_repository

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func getByUsernamePasswordQuery() string {
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
		username = $1 and password = $2
    `
}

// GetByUsernamePassword returns a single row of User by username, password
func (r *repository) GetByUsernamePassword(ctx context.Context, un, pass string) (*User, error) {
	// hash password
	hashed, err := hash(pass)
	if err != nil {
		log.Err(errors.Wrap(err, _errHashingPassword.Error()))
		return &User{}, errors.Wrap(err, _errHashingPassword.Error())
	}

	// execute query
	query := getByUsernamePasswordQuery()
	row := r.DB.DB.QueryRowContext(ctx, query, un, hashed)

	// parse result
	var id int64
	var username string
	var email string
	var createdOn int64
	var lastLogin int64

	err = row.Scan(
		&id, &username, &email, &createdOn, &lastLogin,
	)
	if err != nil {
		log.Err(errors.Wrap(err, _errParsingRowFromDatabase.Error()))
		return &User{}, errors.Wrap(err, _errParsingRowFromDatabase.Error())
	}

	user := &User{
		Id:        id,
		Username:  username,
		Email:     email,
		CreatedOn: createdOn,
		LastLogin: lastLogin,
	}

	return user, nil
}
