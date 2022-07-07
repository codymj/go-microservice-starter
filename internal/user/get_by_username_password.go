package user

import (
	"context"
	"github.com/rs/zerolog/log"
	"go-microservice-starter/internal/repository/user_repository"
)

// GetByUsernamePassword returns a single user by username, password
func (s *service) GetByUsernamePassword(ctx context.Context, un, pass string) (user_repository.User, error) {
	// log info
	log.Info().
		Str("username", un).
		Str("password", pass).
		Msg("user:GetById")

	// get all users via repository
	user, err := s.ur.GetByUsernamePassword(ctx, un, pass)
	if err != nil {
		return user_repository.User{}, err
	}

	return user, nil
}
