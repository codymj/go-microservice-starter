package password

import (
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

// CompareHash is used to compare a user-inputed password to a hash
func (*service) CompareHash(password, hash string) (bool, error) {
	// split hash into parts
	parts := strings.Split(hash, "$")

	// scan params
	cfg := &Config{}
	_, err := fmt.Sscanf(
		parts[3],
		"m=%d,t=%d,p=%d",
		&cfg.Memory,
		&cfg.Time,
		&cfg.Threads,
	)
	if err != nil {
		return false, err
	}

	// extract salt
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	// extract decoded hash and length
	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}
	cfg.KeyLen = uint32(len(decodedHash))

	// hash input to compare
	comparisonHash := argon2.IDKey(
		[]byte(password),
		salt,
		cfg.Time,
		cfg.Memory,
		cfg.Threads,
		cfg.KeyLen,
	)

	return subtle.ConstantTimeCompare(decodedHash, comparisonHash) == 1, nil
}
