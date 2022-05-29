package health

import "context"

// service dependencies to inject
type service struct{}

// Service contract
type Service interface {
	CheckHealth(ctx context.Context) error
}

// New returns an initialized instance
func New() service {
	return service{}
}
