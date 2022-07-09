package users

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"go-microservice-starter/internal/users/users_dao"
)

// GetById returns a single users by id
func (s *service) GetById(ctx context.Context, id uuid.UUID) (*users_dao.User, error) {
	// log info
	log.Info().
		Str("id", id.String()).
		Msg("users:GetById")

	// get users via repository
	user, err := s.ur.GetById(ctx, id)
	if err != nil {
		return &users_dao.User{}, err
	}

	return user, nil
}
