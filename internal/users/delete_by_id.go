package users

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// DeleteById a users by id
func (s *service) DeleteById(ctx context.Context, id uuid.UUID) error {
	// log info
	log.Info().
		Str("id", id.String()).
		Msg(InfoBeginDeleteUserById)

	// get all users via repository
	err := s.ur.DeleteById(ctx, id)
	if err != nil {
		log.Err(err)
		return err
	}

	// log info
	log.Info().
		Str("id", id.String()).
		Msg(InfoEndDeleteUserById)

	return nil
}
