package lexer

import "unicode"

type TokenType int

const (
	TOKEN_LEFT_BRACE TokenType = iota
	TOKEN_RIGHT_BRACE
	TOKEN_EOF
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	input        []byte
	position     int
	readPosition int
	ch           byte
	size         int
}

func NewLexer(input []byte) *Lexer {
	lxr := &Lexer{input: input}
	lxr.size = len(input)
	lxr.nextChar()
	return lxr
}

func (lxr *Lexer) nextChar() {
	if lxr.readPosition >= lxr.size {
		lxr.ch = 0
	} else {
		lxr.ch = lxr.input[lxr.readPosition]
	}
	lxr.position = lxr.readPosition
	lxr.readPosition += 1
}

func (lxr *Lexer) skipWhitespace() {
	for unicode.IsSpace(rune(lxr.ch)) {
		lxr.nextChar()
	}
}

func (lxr *Lexer) NextToken() Token {
	lxr.skipWhitespace()

	var token Token

	switch lxr.ch {
	case '{':
		token = Token{Type: TOKEN_LEFT_BRACE, Value: "{"}
	case '}':
		token = Token{Type: TOKEN_RIGHT_BRACE, Value: "}"}
	default:
		token.Type = TOKEN_EOF
	}

	lxr.nextChar()

	return token
}
