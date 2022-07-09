package password

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestPassword(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "password test suite")
}
