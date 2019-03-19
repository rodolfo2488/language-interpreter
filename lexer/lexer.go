package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
		l.readChar()
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
		l.readChar()
	case '(':
		tok = newToken(token.LPAREN, l.ch)
		l.readChar()
	case ')':
		tok = newToken(token.RPAREN, l.ch)
		l.readChar()
	case ',':
		tok = newToken(token.COMMA, l.ch)
		l.readChar()
	case '+':
		tok = newToken(token.PLUS, l.ch)
		l.readChar()
	case '{':
		tok = newToken(token.LBRACE, l.ch)
		l.readChar()
	case '}':
		tok = newToken(token.RBRACE, l.ch)
		l.readChar()
	case '-':
		tok = newToken(token.MINUS, l.ch)
		l.readChar()
	case '/':
		tok = newToken(token.SLASH, l.ch)
		l.readChar()
	case '*':
		tok = newToken(token.ASTERIKS, l.ch)
		l.readChar()
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
		l.readChar()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		l.readChar()
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookUpIdent(tok.Literal)
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
			l.readPosition = len(l.input)
			l.readChar()
		}
	}

	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(input byte) bool {
	return 'a' <= input && input <= 'z' || 'A' <= input && input <= 'Z' || input == '_'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(input byte) bool {
	return '0' <= input && input <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) Len() int {
	return len(l.input)
}
