package test_utils

import (
	"monkey/lexer"
	"monkey/token"
)

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
