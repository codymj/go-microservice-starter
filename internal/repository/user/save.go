package user

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
    `
}

// Save a new User into the database
func (s *service) Save(ctx context.Context, user User) (User, error) {
	// hash password
	hashed, err := hash(user.Password)
	if err != nil {
		log.Err(errors.Wrap(err, ErrHashingPassword.Error()))
		return User{}, errors.Wrap(err, ErrHashingPassword.Error())
	}
	user.Password = hashed

	// set created_on, last_login
	user.CreatedOn = time.Now().UnixMilli()
	user.LastLogin = user.CreatedOn

	// execute query
	query := saveQuery()
	_, err = s.DB.DB.ExecContext(
		ctx,
		query,
		user.Username,
		user.Password,
		user.Email,
		user.CreatedOn,
		user.LastLogin,
	)
	if err != nil {
		log.Err(errors.Wrap(err, ErrSavingToDatabase.Error()))
		return User{}, errors.Wrap(err, ErrSavingToDatabase.Error())
	}

	return user, nil
}
