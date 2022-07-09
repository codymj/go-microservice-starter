package users_dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"io"
	"strings"
)

var (
	// a valid list of query params
	_validUserParams = []string{
		"id",
		"username",
		"isVerified",
		"createdOn",
		"updatedOn",
	}

	// a map of query param to db field
	_userParamtoDBField = map[string]string{
		"id":         "id",
		"username":   "username",
		"isVerified": "is_verified",
		"createdOn":  "created_on",
		"updatedOn":  "last_login",
	}
)

func getByUsernamePasswordQuery() string {
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
	$1
    `
}

// GetByParams returns a single row of User by query params
func (r *repository) GetByParams(ctx context.Context, params map[string]string) ([]*User, error) {
	// build "where" clause to replace in query
	query := getByUsernamePasswordQuery()
	whereClause := buildWhereClause(params, _validUserParams, _userParamtoDBField)
	if !strings.EqualFold("", whereClause) {
		whereClause = "where " + whereClause
		query = strings.Replace(query, "$1", whereClause, 1)
	} else {
		// if no params were valid, act as get-all
		query = strings.Replace(query, "$1", "", 1)
	}

	// execute query
	rows, err := r.db.DB.QueryContext(ctx, query)
	if err != nil {
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
