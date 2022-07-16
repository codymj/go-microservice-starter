package password

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCompareHash_MatchPasswordWithHash(t *testing.T) {
	// init service
	cfg := &Config{
		Time:    1,
		Memory:  64 * 1024,
		Threads: 4,
		KeyLen:  32,
	}
	s := New(cfg)

	// mock data
	password := "h4sh_m3"
	hash, err := s.HashPassword(password)
	assert.Nil(t, err)

	// invocation
	isMatching, err := s.CompareHash(password, hash)
	assert.Nil(t, err)

	// assertions
	assert.True(t, isMatching)
}

func TestCompareHash_IncorrectPassword(t *testing.T) {
	// init service
	cfg := &Config{
		Time:    1,
		Memory:  64 * 1024,
		Threads: 4,
		KeyLen:  32,
	}
	s := New(cfg)

	// mock data
	correctPassword := "h4sh_m3"
	givenPassword := "1nc0rrect_p4ssw0rd"
	hash, err := s.HashPassword(correctPassword)
	assert.Nil(t, err)

	// invocation
	isMatching, err := s.CompareHash(givenPassword, hash)
	assert.Nil(t, err)

	// assertions
	assert.False(t, isMatching)
}

func TestCompareHash_ErrScanningConfig(t *testing.T) {
	// init service
	cfg := &Config{
		Time:    1,
		Memory:  64 * 1024,
		Threads: 4,
		KeyLen:  32,
	}
	s := New(cfg)

	// mock data
	password := "h4sh_m3"
	hash, err := s.HashPassword(password)
	assert.Nil(t, err)

	// alter hash to force decode to fail
	hash = strings.Replace(hash, "t=", "", 1)

	// invocation
	isMatching, err := s.CompareHash(password, hash)
	assert.NotNil(t, err)

	// assertions
	assert.False(t, isMatching)
}
