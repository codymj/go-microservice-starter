package password

import (
	"encoding/base64"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestHashPassword_Success(t *testing.T) {
	// init service
	cfg := &Config{
		Time:    1,
		Memory:  64 * 1024,
		Threads: 4,
		KeyLen:  32,
	}
	s := New(cfg)

	// mock data
	hash, err := s.HashPassword("h4sh_m3")
	assert.Nil(t, err)

	// parse result
	parsedCfg := &Config{}
	parts := strings.Split(hash, "$")
	_, err = fmt.Sscanf(
		parts[3],
		"m=%d,t=%d,p=%d",
		&parsedCfg.Memory,
		&parsedCfg.Time,
		&parsedCfg.Threads,
	)
	assert.Nil(t, err)

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	assert.Nil(t, err)

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	assert.Nil(t, err)

	parsedCfg.KeyLen = uint32(len(decodedHash))

	// assertions
	assert.Equal(t, cfg.Time, parsedCfg.Time)
	assert.Equal(t, cfg.Memory, parsedCfg.Memory)
	assert.Equal(t, cfg.Threads, parsedCfg.Threads)
	assert.Equal(t, cfg.KeyLen, parsedCfg.KeyLen)
	assert.NotNil(t, salt)
	assert.NotNil(t, hash)
}
