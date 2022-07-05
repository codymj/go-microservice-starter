package user

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
		password,
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
func (s *service) GetByUsernamePassword(ctx context.Context, un, pass string) (User, error) {
	// hash password
	hashed, err := hash(pass)
	if err != nil {
		log.Err(errors.Wrap(err, _errHashingPassword.Error()))
		return User{}, errors.Wrap(err, _errHashingPassword.Error())
	}

	// execute query
	query := getByUsernamePasswordQuery()
	row := s.DB.DB.QueryRowContext(ctx, query, un, hashed)

	// parse result
	user := User{}
	var id int64
	var username string
	var password string
	var email string
	var createdOn int64
	var lastLogin int64

	err = row.Scan(
		&id, &username, &password, &email, &createdOn, &lastLogin,
	)
	if err != nil {
		log.Err(errors.Wrap(err, _errParsingRow.Error()))
		return User{}, errors.Wrap(err, _errParsingRow.Error())
	}

	user = User{
		Id:        id,
		Username:  username,
		Password:  password,
		Email:     email,
		CreatedOn: createdOn,
		LastLogin: lastLogin,
	}

	return user, nil
}
