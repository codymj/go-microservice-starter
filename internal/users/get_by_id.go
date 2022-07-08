package users

import (
	"context"
	"github.com/rs/zerolog/log"
	"go-microservice-starter/internal/users/users_dao"
)

// GetById returns a single users by id
func (s *service) GetById(ctx context.Context, id int64) (*users_dao.User, error) {
	// log info
	log.Info().
		Int64("id", id).
		Msg("users:GetById")

	// get users via repository
	user, err := s.ur.GetById(ctx, id)
	if err != nil {
		return &users_dao.User{}, err
	}

	return user, nil
}
