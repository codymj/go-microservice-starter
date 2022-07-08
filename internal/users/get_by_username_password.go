package users

import (
	"context"
	"github.com/rs/zerolog/log"
	"go-microservice-starter/internal/users/users_dao"
)

// GetByUsernamePassword returns a single users by username, password
func (s *service) GetByUsernamePassword(ctx context.Context, un, pass string) (*users_dao.User, error) {
	// log info
	log.Info().
		Str("username", un).
		Str("password", pass).
		Msg("users:GetById")

	// get all users via repository
	user, err := s.ur.GetByUsernamePassword(ctx, un, pass)
	if err != nil {
		return &users_dao.User{}, err
	}

	return user, nil
}
