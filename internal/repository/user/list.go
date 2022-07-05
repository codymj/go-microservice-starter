package user

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
func (s *service) List(ctx context.Context) ([]User, error) {
	// execute query
	query := listQuery()
	rows, err := s.DB.DB.QueryContext(ctx, query)
	if err != nil {
		log.Err(errors.Wrap(err, ErrQueryingDatabase.Error()))
		return nil, errors.Wrap(err, ErrQueryingDatabase.Error())
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
			log.Err(errors.Wrap(err, ErrParsingRow.Error()))
			return nil, errors.Wrap(err, ErrParsingRow.Error())
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
