package user

import (
	"context"
	"github.com/rs/zerolog/log"
	"go-microservice-starter/internal/repository/user_repository"
)

// GetAll returns all users
func (s *service) GetAll(ctx context.Context) ([]user_repository.User, error) {
	// log info
	log.Info().
		Msg("user:GetAll")

	// get all users via repository
	users, err := s.ur.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
