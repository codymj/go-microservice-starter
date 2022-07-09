package users_dao

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"time"
)

func updateQuery() string {
	return `
	update Users set
		email = $1,
		last_login = $2
	where id = $3
    `
}

// Update an existing User in the database
func (r *repository) Update(ctx context.Context, user *User) (*User, error) {
	user.LastLogin = time.Now().UnixMilli()

	// execute query
	query := updateQuery()
	_, err := r.DB.DB.ExecContext(
		ctx, query,
		user.Email,
		user.LastLogin,
		user.Id,
	)
	if err != nil {
		log.Err(errors.Wrap(err, ErrUpdatingToDatabase.Error()))
		return &User{}, errors.Wrap(err, ErrUpdatingToDatabase.Error())
	}

	// get updated users
	updatedUser, err := r.GetById(ctx, user.Id)
	if err != nil {
		return &User{}, err
	}

	return updatedUser, nil
}
