package ast_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "monkey/ast"
)

var _ = Describe("Ast", func() {

	It("Does something amazing", func() {
		Expect(Program{}).To(BeTrue())
	})
})
