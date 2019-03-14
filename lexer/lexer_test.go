package lexer_test

import (
	"monkey/lexer"
	"monkey/token"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type ExpectedTestCaseToken struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func insert(original []token.Token, position int, value token.Token) []token.Token {
	l := len(original)
	target := original
	if cap(original) == l {
		target = make([]token.Token, l+1, l+10)
		copy(target, original[:position])
	} else {
		target = append(target, token.Token{})
	}
	copy(target[position+1:], original[position:])
	target[position] = value
	return target
}

func LexerToTokenList(lexer lexer.Lexer) []token.Token {
	list := []token.Token{}
	currentValue := token.Token{}
	currentValue = lexer.NextToken()
	for currentValue.Type != token.EOF {
		list = insert(list, len(list), currentValue)
		currentValue = lexer.NextToken()
	}

	return list
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
