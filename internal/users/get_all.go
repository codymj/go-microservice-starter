package users

import (
	"context"
	"github.com/rs/zerolog/log"
	"go-microservice-starter/internal/repository/users_repository"
)

// GetAll returns all users
func (s *service) GetAll(ctx context.Context) ([]*users_repository.User, error) {
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
