package lexer

import "unicode"

type TokenType int

const (
	TOKEN_LEFT_BRACE TokenType = iota
	TOKEN_RIGHT_BRACE
	TOKEN_STRING
	TOKEN_COLON
	TOKEN_COMMA
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

func (lxr *Lexer) readString() string {
	startPosition := lxr.readPosition + 1

	for {
		lxr.nextChar()
		if lxr.ch == '"' || lxr.ch == 0 {
			break
		}
	}

	return string(lxr.input[startPosition:lxr.readPosition])
}

func (lxr *Lexer) NextToken() Token {
	lxr.skipWhitespace()

	var token Token

	switch lxr.ch {
	case '{':
		token = Token{Type: TOKEN_LEFT_BRACE, Value: "{"}
	case '}':
		token = Token{Type: TOKEN_RIGHT_BRACE, Value: "}"}
	case '"':
		value := lxr.readString()
		token = Token{Type: TOKEN_STRING, Value: value}
	case ':':
		token = Token{Type: TOKEN_COLON, Value: ":"}
	case ',':
		token = Token{Type: TOKEN_COMMA, Value: ","}
	default:
		token.Type = TOKEN_EOF
	}

	lxr.nextChar()

	return token
}
