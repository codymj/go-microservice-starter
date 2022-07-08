package users_dao

import (
	"golang.org/x/crypto/bcrypt"
	"io"
)

// hash a plain-text password and returns stringified hash
func hash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

// Close is a wrapper for defer Close() methods
func Close(err *error, c io.Closer) {
	if e := c.Close(); err == nil {
		*err = e
	}
}
