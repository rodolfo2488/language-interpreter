package parser_test

import (
	"monkey/lexer"
	"monkey/parser"

	. "github.com/onsi/gomega"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("Parser", func() {
	var (
		input   string
		l       *lexer.Lexer
		subject *parser.Parser
	)

	Context("Let statements", func() {
		It("returns a program with one statement for each variable", func() {

			input = `let five = 5;let four = 3;`
			l = lexer.New(input)
			subject = parser.New(l)

			Expect(subject.ParseProgram().Statements).To(HaveLen(2))
		})
	})
})
