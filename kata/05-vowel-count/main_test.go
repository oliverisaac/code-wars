// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo></http:>
// Gomega Matcher Library <http://onsi.github.io/gomega></http:>

package kata_test

import (
	. "codewarrior/kata"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Example", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(GetCount("abracadabra")).To(Equal(5))
	})
})
