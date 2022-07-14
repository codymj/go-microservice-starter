package users_dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func getByIdQuery() string {
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
	where
		id = $1
    `
}

// GetById returns a single row of User by id from database
func (r *repository) GetById(ctx context.Context, id uuid.UUID) (User, error) {
	// execute query
	query := getByIdQuery()
	row := r.db.DB.QueryRowContext(ctx, query, id.String())

	// parse result
	var username string
	var email string
	var isVerified bool
	var createdOn int64
	var updatedOn int64

	err := row.Scan(
		&id, &username, &email, &isVerified, &createdOn, &updatedOn,
	)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return User{}, nil
	} else if err != nil {
		return User{}, errors.Wrap(err, ErrParsingRowFromDatabase.Error())
	}

	user := User{
		Id:         id,
		Username:   username,
		Email:      email,
		IsVerified: isVerified,
		CreatedOn:  createdOn,
		UpdatedOn:  updatedOn,
	}

	return user, nil
}
