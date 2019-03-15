package lexer_test

import (
	"monkey/lexer"
	"monkey/token"

	. "monkey/test_utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type ExpectedTestCaseToken struct {
	expectedType    token.TokenType
	expectedLiteral string
}

var _ = Describe("Lexer", func() {

	It("creates map for simple programs", func() {
		input := `
let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);`

		test := []ExpectedTestCaseToken{
			{token.LET, "let"},
			{token.IDENT, "five"},
			{token.ASSIGN, "="},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "ten"},
			{token.ASSIGN, "="},
			{token.INT, "10"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "add"},
			{token.ASSIGN, "="},
			{token.FUNCTION, "fn"},
			{token.LPAREN, "("},
			{token.IDENT, "x"},
			{token.COMMA, ","},
			{token.IDENT, "y"},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.IDENT, "x"},
			{token.PLUS, "+"},
			{token.IDENT, "y"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "result"},
			{token.ASSIGN, "="},
			{token.IDENT, "add"},
			{token.LPAREN, "("},
			{token.IDENT, "five"},
			{token.COMMA, ","},
			{token.IDENT, "ten"},
			{token.RPAREN, ")"},
			{token.SEMICOLON, ";"},
			{token.EOF, ""},
		}

		l := lexer.New(input)

		for _, tt := range test {
			tok := l.NextToken()

			Expect(tok.Type).To(Equal(tt.expectedType))
			Expect(tok.Literal).To(Equal(tt.expectedLiteral))
		}
	})

	It("parses basic math operations with + - / *", func() {
		input := `let foo=5 - 1 * (5+2/10);`

		subject := lexer.New(input)

		plusToken := token.Token{Type: token.PLUS, Literal: "+"}
		Expect(LexerToTokenList(*subject)).To(ContainElement(plusToken))

		minusToken := token.Token{Type: token.MINUS, Literal: "-"}
		Expect(LexerToTokenList(*subject)).To(ContainElement(minusToken))

		divisionToken := token.Token{Type: token.SLASH, Literal: "/"}
		Expect(LexerToTokenList(*subject)).To(ContainElement(divisionToken))

		asteriskToken := token.Token{Type: token.ASTERIKS, Literal: "*"}
		Expect(LexerToTokenList(*subject)).To(ContainElement(asteriskToken))
	})

	It("parses conditionals", func() {
		input := `if (!foo)  { foo++; } else { foo--; }`

		subject := lexer.New(input)
		IfToken := token.Token{Type: token.IF, Literal: "if"}
		Expect(LexerToTokenList(*subject)).To(ContainElement(IfToken))

		elseToken := token.Token{Type: token.ELSE, Literal: "else"}
		Expect(LexerToTokenList(*subject)).To(ContainElement(elseToken))

		bangToken := token.Token{Type: token.BANG, Literal: "!"}
		Expect(LexerToTokenList(*subject)).To(ContainElement(bangToken))
	})

	It("parses return statements", func() {
		input := `if (foo) return 5; `

		subject := lexer.New(input)
		returnToken := token.Token{Type: token.RETURN, Literal: "return"}
		Expect(LexerToTokenList(*subject)).To(ContainElement(returnToken))
	})

	It("parses true false as keywords", func() {
		input := `foo=true;bar=false;`

		subject := lexer.New(input)
		trueToken := token.Token{Type: token.TRUE, Literal: "true"}
		Expect(LexerToTokenList(*subject)).To(ContainElement(trueToken))

		falseToken := token.Token{Type: token.FALSE, Literal: "false"}
		Expect(LexerToTokenList(*subject)).To(ContainElement(falseToken))
	})

	It("parses two character Operators `==` and `!=`", func() {
		input := `if(bar == foo) {} else if (5 != foo) {}`

		subject := lexer.New(input)
		eqToken := token.Token{Type: token.EQ, Literal: "=="}
		Expect(LexerToTokenList(*subject)).To(ContainElement(eqToken))

		notEqToken := token.Token{Type: token.NOT_EQ, Literal: "!="}
		Expect(LexerToTokenList(*subject)).To(ContainElement(notEqToken))
	})

	Context("when there is an illegal caracter in the program", func() {
		var (
			subject *lexer.Lexer
		)

		BeforeEach(func() {
			input := `let yu^2=5;`

			subject = lexer.New(input)
		})

		It("parses the illegal char", func() {
			expectedIllegalToken := token.Token{Type: token.ILLEGAL, Literal: "^"}
			Expect(LexerToTokenList(*subject)).To(ContainElement(expectedIllegalToken))
		})

		It("stop the parse at the illegal char", func() {
			Expect(LexerToTokenList(*subject)).Should(HaveLen(3))
		})
	})
})
