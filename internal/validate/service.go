package validate

import (
	"context"
)

// service dependencies to inject
type service struct{}

// Service contract
type Service interface {
	ValidatePostUsers(ctx context.Context, body []byte) ([]string, error)
}

// New returns an initialized instance
func New() Service {
	return &service{}
}
