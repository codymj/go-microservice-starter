package user

import (
	"context"
	"github.com/rs/zerolog/log"
	"go-microservice-starter/internal/repository/user_repository"
)

// Update an existing user
func (s *service) Update(ctx context.Context, r PutUserRequest) (*user_repository.User, error) {
	// log info
	log.Info().
		Interface("request", r).
		Msg("user:Create")

	// transform
	unupdatedUser := transformPutUserRequest(r)

	// save user via repository
	updatedUser, err := s.ur.Save(ctx, &unupdatedUser)
	if err != nil {
		return &user_repository.User{}, err
	}

	return updatedUser, nil
}
