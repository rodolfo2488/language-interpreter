package lexer

import (
	"monkey/token"
	"testing"
)

type ExpectedTestCaseToken struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestNextToken(t *testing.T) {

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

	assertTokenMap(t, test, input)
}

func TestDontCareSpacing(t *testing.T) {
	input := `    let        five        =            5;       `

	tests := []ExpectedTestCaseToken{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	assertTokenMap(t, tests, input)
}

func TestIllegalChars(t *testing.T) {

	input := `let yu^2= 5;`

	tests := []ExpectedTestCaseToken{
		{token.LET, "let"},
		{token.IDENT, "yu"},
		{token.ILLEGAL, "^"},
	}

	assertTokenMap(t, tests, input)
}
func assertTokenMap(t *testing.T, expectedOutput []ExpectedTestCaseToken, input string) {
	l := New(input)

	for i, tt := range expectedOutput {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] - tokenliteral wrong. expecte=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
