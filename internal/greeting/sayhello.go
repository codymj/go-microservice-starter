package greeting

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
)

// SayHello is just an example of some business logic
func (s *service) SayHello(_ context.Context, r PostGreetingRequest) PostGreetingResponse {
	// log info
	log.Info().
		Interface("request", r).
		Msg("greeting:sayhello")

	return PostGreetingResponse{
		Status:  "ok",
		Message: fmt.Sprintf("hello, %s!", r.Name),
	}
}
