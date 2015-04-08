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

	Context("#UniqueInts", func() {

		It("makes the slice unique", func() {
			i := []string{1, 2, 3, 2, 4, 1, 5, 6}
			Expect(ormtools.UniqueInts(i)).To(Equal([]string{1, 2, 3, 4, 5, 6}))
		})

	})

})
