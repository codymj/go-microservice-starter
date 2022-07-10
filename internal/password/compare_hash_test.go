package password_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go-microservice-starter/internal/password"
)

var _ = Describe("CompareHash()", func() {
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

	It("should match a password with its hash", func() {
		// hash to compare
		password := "h4sh_m3"
		hash, err := s.HashPassword(password)
		Expect(err).To(BeNil())

		// invocation
		isMatching, err := s.CompareHash(password, hash)
		Expect(err).To(BeNil())

		// assertions
		Expect(isMatching).To(BeTrue())
	})

	It("should not match an incorrect password", func() {
		// hashes to compare
		correctPassword := "h4sh_m3"
		givenPassword := "1nc0rrect_p4ssw0rd"
		hash, err := s.HashPassword(correctPassword)
		Expect(err).To(BeNil())

		// invocation
		isMatching, err := s.CompareHash(givenPassword, hash)
		Expect(err).To(BeNil())

		// assertions
		Expect(isMatching).To(BeFalse())
	})
})
