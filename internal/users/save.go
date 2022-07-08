package users

import (
	"context"
	"github.com/rs/zerolog/log"
	"go-microservice-starter/internal/repository/users_repository"
)

// Save creates a new users
func (s *service) Save(ctx context.Context, r PostUsersRequest) (*users_repository.User, error) {
	// log info
	log.Info().
		Interface("request", r).
		Msg("users:Create")

	// transform
	unsavedUser := transformPostUserRequest(r)

	// save users via repository
	savedUser, err := s.ur.Save(ctx, &unsavedUser)
	if err != nil {
		return &users_repository.User{}, err
	}

	return savedUser, nil
}
