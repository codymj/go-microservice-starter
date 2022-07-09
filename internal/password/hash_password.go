package password

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
)

// HashPassword generates a password hash
func (s *service) HashPassword(password string) (string, error) {
	// generate salt
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	// hash
	hash := argon2.IDKey(
		[]byte(password),
		salt,
		s.cfg.Time,
		s.cfg.Memory,
		s.cfg.Threads,
		s.cfg.KeyLen,
	)

	// base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// format full password hash
	format := "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
	full := fmt.Sprintf(
		format,
		argon2.Version,
		s.cfg.Memory,
		s.cfg.Time,
		s.cfg.Threads,
		b64Salt,
		b64Hash,
	)

	return full, nil
}
