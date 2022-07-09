package users

import (
	"context"
	"github.com/rs/zerolog/log"
	"go-microservice-starter/internal/users/users_dao"
)

// Update an existing users
func (s *service) Update(ctx context.Context, id int64, r PutUsersIdRequest) (*users_dao.User, error) {
	// log info
	log.Info().
		Int64("id", id).
		Interface("request", r).
		Msg("users:Update")

	// get user to update
	nonupdatedUser, err := s.ur.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	// override updatable fields
	nonupdatedUser.Email = r.Email

	// save users via repository
	updatedUser, err := s.ur.Update(ctx, nonupdatedUser)
	if err != nil {
		return &users_dao.User{}, err
	}

	return updatedUser, nil
}
