package test_publicPosts_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTestPublicPosts(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestPublicPosts Suite")
}
