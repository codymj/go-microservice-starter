package validate

import (
	"context"
	"github.com/codymj/jsonvalidator/jsonvalidator"
	"github.com/rs/zerolog/log"
)

// ValidatePostGreeting validates the payload to POST /greeting endpoint
func (s *service) ValidatePostGreeting(_ context.Context, payload []byte) ([]string, error) {
	// log info
	log.Info().
		RawJSON("payload", payload).
		Msg("validate:validatepostgreeting")

	// validate payload against schema
	errors, err := jsonvalidator.Validate(getPostGreetingSchema(), payload)
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
