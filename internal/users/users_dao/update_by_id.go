package users_dao

import (
	"context"
	"github.com/pkg/errors"
	"time"
)

func updateQuery() string {
	return `
	update Users set
		email = $1,
		is_verified = $2,
		updated_on = $3
	where id = $4
    `
}

// UpdateById an existing User in the database by id
func (r *repository) UpdateById(ctx context.Context, user User) (User, error) {
	user.UpdatedOn = time.Now().UnixMilli()

	// execute query
	query := updateQuery()
	_, err := r.db.DB.ExecContext(
		ctx, query,
		user.Email,
		user.IsVerified,
		user.UpdatedOn,
		user.Id,
	)
	if err != nil {
		return User{}, errors.Wrap(err, ErrUpdatingToDatabase.Error())
	}

	// get updated users
	updatedUser, err := r.GetById(ctx, user.Id)
	if err != nil {
		return User{}, err
	}

	return updatedUser, nil
}
