package users

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"go-microservice-starter/internal/users/users_dao"
)

// UpdateById an existing user by id
func (s *service) UpdateById(ctx context.Context, id uuid.UUID, r PutUsersIdRequest) (users_dao.User, error) {
	// log info
	log.Info().
		Str("id", id.String()).
		Interface("request", r).
		Msg(InfoBeginUpdateUserById)

	// get user to update
	userToUpdate, err := s.ur.GetById(ctx, id)
	if err != nil {
		log.Err(err)
		return users_dao.User{}, err
	}

	// override updatable fields
	userToUpdate.Email = r.Email
	userToUpdate.IsVerified = r.IsVerified

	// update user via repository
	updatedUser, err := s.ur.UpdateById(ctx, userToUpdate)
	if err != nil {
		log.Err(err)
		return users_dao.User{}, err
	}

	// log info
	log.Info().
		Str("id", id.String()).
		Interface("request", r).
		Msg(InfoEndUpdateUserById)

	return updatedUser, nil
}
