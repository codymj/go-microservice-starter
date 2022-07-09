package password

// service dependencies to inject
type service struct {
	cfg *Config
}

// Service contract
type Service interface {
	HashPassword(password string) (string, error)
	CompareHash(password, hash string) (bool, error)
}

// New returns an initialized instance
func New(cfg *Config) Service {
	return &service{
		cfg: cfg,
	}
}
