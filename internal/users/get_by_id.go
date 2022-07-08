package users

import (
	"context"
	"github.com/rs/zerolog/log"
	"go-microservice-starter/internal/repository/user_repository"
)

// GetById returns a single users by id
func (s *service) GetById(ctx context.Context, id int64) (*user_repository.User, error) {
	// log info
	log.Info().
		Int64("id", id).
		Msg("users:GetById")

	// get users via repository
	user, err := s.ur.GetById(ctx, id)
	if err != nil {
		return &user_repository.User{}, err
	}

	return user, nil
}
