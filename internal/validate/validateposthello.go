package validate

import (
	"context"
	"github.com/codymj/jsonvalidator/jsonvalidator"
	"github.com/rs/zerolog/log"
)

// ValidatePostHello validates the payload to post /hello endpoint
func (s *service) ValidatePostHello(_ context.Context, payload []byte) ([]string, error) {
	// log info
	log.Info().
		RawJSON("payload", payload).
		Msg("validating POST /hello payload")

	// validate payload against schema
	errors, err := jsonvalidator.Validate(getPostHelloSchema(), payload)
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
