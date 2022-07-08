package users_dao

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
	"strings"
)

var (
	// a valid list of query params
	_validParams = []string{
		"username",
		"createdOn",
		"lastLogin",
	}

	// a map of query param to db field
	_queryParamtoDatbaseField = map[string]string{
		"username":  "username",
		"createdOn": "created_on",
		"lastLogin": "last_login",
	}
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
	$1
    `
}

func buildWhereClause(params map[string]string) string {
	clauses := make([]string, 0)
	for _, validParam := range _validParams {
		_, ok := params[validParam]
		if ok {
			databaseField := _queryParamtoDatbaseField[validParam]
			value := params[validParam]
			clause := fmt.Sprintf("%s='%s'", databaseField, value)
			clauses = append(clauses, clause)
		}
	}

	return strings.Join(clauses, " and ")
}

// GetByParams returns a single row of User by query params
func (r *repository) GetByParams(ctx context.Context, params map[string]string) ([]*User, error) {
	// build where clause to replace in query
	query := getByUsernamePasswordQuery()
	whereClause := buildWhereClause(params)
	if !strings.EqualFold("", whereClause) {
		whereClause = "where " + whereClause
		query = strings.Replace(query, "$1", whereClause, 1)
	} else {
		query = strings.Replace(query, "$1", "", 1)
	}
	fmt.Println(query)

	// execute query
	rows, err := r.DB.DB.QueryContext(ctx, query)
	if err != nil {
		log.Err(errors.Wrap(err, _errQueryingDatabase.Error()))
		return nil, errors.Wrap(err, _errQueryingDatabase.Error())
	}
	defer Close(&err, io.Closer(rows))

	// parse result
	users := make([]*User, 0)
	for rows.Next() {
		var id int64
		var username string
		var email string
		var createdOn int64
		var lastLogin int64

		err = rows.Scan(
			&id, &username, &email, &createdOn, &lastLogin,
		)
		if err != nil {
			log.Err(errors.Wrap(err, _errParsingRowFromDatabase.Error()))
			return nil, errors.Wrap(err, _errParsingRowFromDatabase.Error())
		}

		user := User{
			Id:        id,
			Username:  username,
			Email:     email,
			CreatedOn: createdOn,
			LastLogin: lastLogin,
		}
		users = append(users, &user)
	}

	return users, nil
}
