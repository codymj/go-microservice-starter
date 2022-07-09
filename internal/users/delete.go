package users

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// Delete a users by id
func (s *service) Delete(ctx context.Context, id uuid.UUID) error {
	// log info
	log.Info().
		Str("id", id.String()).
		Msg("users:GetById")

	// get all users via repository
	err := s.ur.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
