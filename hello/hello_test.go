package hello_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/elgnay/travis-learning/hello"
)

const message = "Hello, world."

var _ = Describe("Hello", func() {
	It("Should output hello world", func() {
		Expect(Hello()).ShouldNot(BeZero())
		Expect(Hello()).Should(Equal(message))
	})
})
