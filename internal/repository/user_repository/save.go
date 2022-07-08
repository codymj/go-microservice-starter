package user_repository

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"time"
)

func saveQuery() string {
	return `
	insert into Users (
		username,
		password,
		email,
		created_on,
		last_login
	)
	values (
		$1, $2, $3, $4, $5
	)
	returning id
    `
}

// Save a new User into the database
func (r *repository) Save(ctx context.Context, user *User) (*User, error) {
	// hash password
	hashed, err := hash(user.Password)
	if err != nil {
		log.Err(errors.Wrap(err, _errHashingPassword.Error()))
		return &User{}, errors.Wrap(err, _errHashingPassword.Error())
	}
	user.Password = hashed
	user.CreatedOn = time.Now().UnixMilli()
	user.LastLogin = user.CreatedOn

	// execute query
	var lastInsertedId int64
	query := saveQuery()
	row := r.DB.DB.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Password,
		user.Email,
		user.CreatedOn,
		user.LastLogin,
	)
	err = row.Scan(&lastInsertedId)
	if err != nil {
		log.Err(errors.Wrap(err, _errSavingToDatabase.Error()))
		return &User{}, errors.Wrap(err, _errSavingToDatabase.Error())
	}

	// get saved users
	savedUser, err := r.GetById(ctx, lastInsertedId)
	if err != nil {
		return &User{}, err
	}

	return savedUser, nil
}
