package users_handler_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestUsersHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "users handler test suite")
}
