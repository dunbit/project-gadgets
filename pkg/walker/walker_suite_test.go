package walker_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWalker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Walker Suite")
}
