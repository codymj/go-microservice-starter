package users

import (
	"context"
	"github.com/rs/zerolog/log"
	"go-microservice-starter/internal/repository/users_repository"
)

// GetByUsernamePassword returns a single users by username, password
func (s *service) GetByUsernamePassword(ctx context.Context, un, pass string) (*users_repository.User, error) {
	// log info
	log.Info().
		Str("username", un).
		Str("password", pass).
		Msg("users:GetById")

	// get all users via repository
	user, err := s.ur.GetByUsernamePassword(ctx, un, pass)
	if err != nil {
		return &users_repository.User{}, err
	}

	return user, nil
}
