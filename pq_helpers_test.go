package ormtools_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/oreillymedia/go-ormtools"
)

var _ = Describe("Postgres helpers", func() {

	Context("#FromPqStringArray", func() {

		It("works on string with comma spaces", func() {
			arr := "{rune, steve, zach, fred, martin}"
			Expect(ormtools.FromPqStringArray(arr)).To(Equal([]string{"rune", "steve", "zach", "fred", "martin"}))
		})

		FIt("works on empty array", func() {
			arr := "{}"
			Expect(len(ormtools.FromPqStringArray(arr))).To(Equal(0))
		})

	})

})
