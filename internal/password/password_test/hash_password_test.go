package password_test

import (
	"encoding/base64"
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go-microservice-starter/internal/password"
	"strings"
)

var _ = Describe("HashPassword()", func() {
	// stub service
	var cfg *password.Config
	var s password.Service

	BeforeEach(func() {
		cfg = &password.Config{
			Time:    1,
			Memory:  64 * 1024,
			Threads: 4,
			KeyLen:  32,
		}
		s = password.New(cfg)
	})

	It("should hash a given password", func() {
		// invocation
		hash, err := s.HashPassword("h4sh_m3")
		Expect(err).To(BeNil())

		// parse hash
		parsedCfg := &password.Config{}
		parts := strings.Split(hash, "$")
		_, err = fmt.Sscanf(
			parts[3],
			"m=%d,t=%d,p=%d",
			&parsedCfg.Memory,
			&parsedCfg.Time,
			&parsedCfg.Threads,
		)
		Expect(err).To(BeNil())

		salt, err := base64.RawStdEncoding.DecodeString(parts[4])
		Expect(err).To(BeNil())

		decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
		Expect(err).To(BeNil())
		parsedCfg.KeyLen = uint32(len(decodedHash))

		// assertions
		Expect(parsedCfg.Time).To(Equal(cfg.Time))
		Expect(parsedCfg.Memory).To(Equal(cfg.Memory))
		Expect(parsedCfg.Threads).To(Equal(cfg.Threads))
		Expect(parsedCfg.KeyLen).To(Equal(cfg.KeyLen))
		Expect(salt).To(Not(BeNil()))
		Expect(hash).To(Not(BeNil()))
	})
})
