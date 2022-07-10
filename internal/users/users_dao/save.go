package users_dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

func saveQuery() string {
	return `
	insert into Users (
		id,
		username,
		password,
		email,
		is_verified,
		created_on,
		updated_on
	)
	values (
		$1, $2, $3, $4, $5, $6, $7
	)
    `
}

// Save a new User into the database
func (r *repository) Save(ctx context.Context, user *User) (*User, error) {
	// hash password
	hashed, err := r.ps.HashPassword(user.Password)
	if err != nil {
		return &User{}, errors.Wrap(err, ErrHashingPassword.Error())
	}
	user.Id = uuid.New()
	user.Password = hashed
	user.CreatedOn = time.Now().UnixMilli()
	user.UpdatedOn = user.CreatedOn

	// execute query
	query := saveQuery()
	_, err = r.db.DB.ExecContext(
		ctx,
		query,
		user.Id.String(),
		user.Username,
		user.Password,
		user.Email,
		user.IsVerified,
		user.CreatedOn,
		user.UpdatedOn,
	)
	if err != nil {
		return &User{}, errors.Wrap(err, ErrSavingToDatabase.Error())
	}

	// don't return password hash in response
	user.Password = ""

	return user, nil
}
