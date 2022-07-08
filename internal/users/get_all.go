package users

import (
	"context"
	"github.com/rs/zerolog/log"
	"go-microservice-starter/internal/users/users_dao"
)

// GetAll returns all users
func (s *service) GetAll(ctx context.Context) ([]*users_dao.User, error) {
	// log info
	log.Info().
		Msg("users:GetAll")

	// get all users via repository
	users, err := s.ur.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
