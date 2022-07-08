package user_repository

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"time"
)

func updateQuery() string {
	return `
	update Users set
		password = $1,
		email = $2,
		last_login = $3
	)
	where id = $4
    `
}

// Update an existing User in the database
func (r *repository) Update(ctx context.Context, user *User) (*User, error) {
	// hash password
	hashed, err := hash(user.Password)
	if err != nil {
		log.Err(errors.Wrap(err, _errHashingPassword.Error()))
		return &User{}, errors.Wrap(err, _errHashingPassword.Error())
	}
	user.Password = hashed
	user.LastLogin = time.Now().UnixMilli()

	// execute query
	query := updateQuery()
	_, err = r.DB.DB.ExecContext(
		ctx, query,
		user.Password,
		user.Email,
		user.LastLogin,
		user.Id,
	)
	if err != nil {
		log.Err(errors.Wrap(err, _errUpdatingToDatabase.Error()))
		return &User{}, errors.Wrap(err, _errUpdatingToDatabase.Error())
	}

	// get updated users
	updatedUser, err := r.GetById(ctx, user.Id)
	if err != nil {
		return &User{}, err
	}

	return updatedUser, nil
}
