package users

import (
	"context"
	"github.com/rs/zerolog/log"
	"go-microservice-starter/internal/users/users_dao"
)

// Save creates a new users
func (s *service) Save(ctx context.Context, r PostUsersRequest) (*users_dao.User, error) {
	// log info
	log.Info().
		Interface("request", r).
		Msg("users:Create")

	// transform
	unsavedUser := postUsersRequestToDAO(r)

	// save users via repository
	savedUser, err := s.ur.Save(ctx, &unsavedUser)
	if err != nil {
		return &users_dao.User{}, err
	}

	return savedUser, nil
}
