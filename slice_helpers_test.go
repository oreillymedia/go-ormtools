package ormtools_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/oreillymedia/go-ormtools"
)

var _ = Describe("Slice helpers", func() {

	Context("#UniqueStrings", func() {

		It("makes the slice unique", func() {
			s := []string{"a", "b", "c", "b", "d", "e", "a", "e"}
			Expect(ormtools.UniqueStrings(s)).To(Equal([]string{"a", "b", "c", "d", "e"}))
		})

	})

})
