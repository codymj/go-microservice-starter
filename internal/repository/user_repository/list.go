package user_repository

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
)

func listQuery() string {
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
    `
}

// List returns all rows of User from database
func (r *repository) List(ctx context.Context) ([]User, error) {
	// execute query
	query := listQuery()
	rows, err := r.DB.DB.QueryContext(ctx, query)
	if err != nil {
		log.Err(errors.Wrap(err, _errQueryingDatabase.Error()))
		return nil, errors.Wrap(err, _errQueryingDatabase.Error())
	}
	defer Close(&err, io.Closer(rows))

	// parse result
	users := make([]User, 0)
	for rows.Next() {
		var id int64
		var username string
		var password string
		var email string
		var createdOn int64
		var lastLogin int64

		err = rows.Scan(
			&id, &username, &password, &email, &createdOn, &lastLogin,
		)
		if err != nil {
			log.Err(errors.Wrap(err, _errParsingRowFromDatabase.Error()))
			return nil, errors.Wrap(err, _errParsingRowFromDatabase.Error())
		}

		user := User{
			Id:        id,
			Username:  username,
			Password:  password,
			Email:     email,
			CreatedOn: createdOn,
			LastLogin: lastLogin,
		}
		users = append(users, user)
	}

	return users, nil
}
