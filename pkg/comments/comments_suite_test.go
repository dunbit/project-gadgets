package comments_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestComments(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Comments Suite")
}
