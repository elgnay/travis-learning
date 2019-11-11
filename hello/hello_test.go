package hello_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/elgnay/travis-learning/hello"
)

var _ = Describe("Hello", func() {
	It("Should output hello world", func() {
		Expect(Hello()).Should(Equal("Hello, world."))
	})
})
