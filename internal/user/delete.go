package user

import (
	"context"
	"github.com/rs/zerolog/log"
)

// Delete a user by id
func (s *service) Delete(ctx context.Context, id int64) error {
	// log info
	log.Info().
		Int64("id", id).
		Msg("user:GetById")

	// get all users via repository
	err := s.ur.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
