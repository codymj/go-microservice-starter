package user

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
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
func (s *service) Update(ctx context.Context, user User) (User, error) {
	// hash password
	hashed, err := hash(user.Password)
	if err != nil {
		log.Err(errors.Wrap(err, ErrHashingPassword.Error()))
		return User{}, errors.Wrap(err, ErrHashingPassword.Error())
	}
	user.Password = hashed

	// execute query
	query := updateQuery()
	_, err = s.DB.DB.ExecContext(
		ctx, query,
		user.Password,
		user.Email,
		user.LastLogin,
		user.Id,
	)
	if err != nil {
		log.Err(errors.Wrap(err, ErrUpdatingToDatabase.Error()))
		return User{}, errors.Wrap(err, ErrUpdatingToDatabase.Error())
	}

	return user, nil
}
