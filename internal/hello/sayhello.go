package hello

import (
	"context"
	"fmt"
)

// SayHello is just an example of some business logic
func (s *service) SayHello(_ context.Context, r PostRequest) PostResponse {
	return PostResponse{
		Status:  "ok",
		Message: fmt.Sprintf("hello, %s", r.Name),
	}
}
