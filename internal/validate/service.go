package validate

import (
	"context"
)

// service dependencies to inject
type service struct{}

// Service contract
type Service interface {
	ValidatePostGreeting(ctx context.Context, payload []byte) ([]string, error)
}

// New returns an initialized instance
func New() Service {
	return &service{}
}
