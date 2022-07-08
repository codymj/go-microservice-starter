package users

import (
	"context"
	"github.com/rs/zerolog/log"
	"go-microservice-starter/internal/repository/user_repository"
)

// Save creates a new users
func (s *service) Save(ctx context.Context, r PostUserRequest) (*user_repository.User, error) {
	// log info
	log.Info().
		Interface("request", r).
		Msg("users:Create")

	// transform
	unsavedUser := transformPostUserRequest(r)

	// save users via repository
	savedUser, err := s.ur.Save(ctx, &unsavedUser)
	if err != nil {
		return &user_repository.User{}, err
	}

	return savedUser, nil
}
