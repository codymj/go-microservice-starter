package users

import (
	"context"
	"github.com/rs/zerolog/log"
	"go-microservice-starter/internal/users/users_dao"
)

// GetByParams returns a single users by query params
func (s *service) GetByParams(ctx context.Context, params map[string]string) ([]*users_dao.User, error) {
	// log info
	log.Info().
		Interface("params", params).
		Msg("users:GetById")

	// get all users via repository
	users, err := s.ur.GetByParams(ctx, params)
	if err != nil {
		return nil, err
	}

	return users, nil
}