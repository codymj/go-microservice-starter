package validate

import (
	"context"
	"github.com/codymj/jsonvalidator/jsonvalidator"
	"github.com/rs/zerolog/log"
)

// ValidatePostUsers validates the payload to POST /users endpoint
func (s *service) ValidatePostUsers(_ context.Context, body []byte) ([]string, error) {
	// compact json request for logging
	compacted, err := compactJson(body)
	if err != nil {
		return nil, err
	}

	// log info
	log.Info().
		RawJSON("payload", compacted).
		Msg("validate:ValidatePostUsers")

	// validate payload against schema
	errors, err := jsonvalidator.Validate(getPostUsersSchema(), compacted)
	if err != nil {
		return nil, err
	}

	// check errors
	if len(errors) > 0 {
		return errors, nil
	} else {
		return nil, nil
	}
}
