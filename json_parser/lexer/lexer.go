package lexer

import (
	"io"
	"unicode"
)

type TokenType int

const (
	TOKEN_LEFT_BRACE TokenType = iota
	TOKEN_RIGHT_BRACE
	TOKEN_LEFT_BRACKET
	TOKEN_RIGHT_BRACKET
	TOKEN_STRING
	TOKEN_COLON
	TOKEN_COMMA
	TOKEN_NUMBER
	TOKEN_TRUE
	TOKEN_FALSE
	TOKEN_NULL
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

func NewLexer(r io.Reader) *Lexer {
	input, err := io.ReadAll(r)
	if err != nil {
		panic("failed to read input")
	}
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

func (lxr *Lexer) advancePosition(n int) {
	for i := 0; i < n; i++ {
		lxr.nextChar()
	}
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

func (lxr *Lexer) readNumber() string {
	startPosition := lxr.position

	for unicode.IsDigit(rune(lxr.ch)) || lxr.ch == '-' {
		lxr.nextChar()
	}

	return string(lxr.input[startPosition:lxr.position])
}

func (lxr *Lexer) readBoolean() (TokenType, string) {
	if lxr.matchWord("true") {
		return TOKEN_TRUE, "true"
	} else if lxr.matchWord("false") {
		return TOKEN_FALSE, "false"
	}

	return TOKEN_EOF, ""
}

func (lxr *Lexer) readNull() string {
	if lxr.matchWord("null") {
		return "null"
	}
	return ""
}

func (lxr *Lexer) matchWord(word string) bool {
	if lxr.size-lxr.position < len(word) {
		return false
	}

	if string(lxr.input[lxr.position:lxr.position+len(word)]) != word {
		return false
	}

	nextChar := lxr.input[lxr.position+len(word)]
	// next character can't be a valid identifier char, like "trueK" or "falseee"
	if unicode.IsLetter(rune(nextChar)) || unicode.IsDigit(rune(nextChar)) || nextChar == '_' {
		return false
	}

	lxr.advancePosition(len(word))

	return true
}

func (lxr *Lexer) NextToken() Token {
	lxr.skipWhitespace()

	var token Token

	switch lxr.ch {
	case '{':
		token = Token{Type: TOKEN_LEFT_BRACE, Value: "{"}
	case '}':
		token = Token{Type: TOKEN_RIGHT_BRACE, Value: "}"}
	case '[':
		token = Token{Type: TOKEN_LEFT_BRACKET, Value: "["}
	case ']':
		token = Token{Type: TOKEN_RIGHT_BRACKET, Value: "]"}
	case '"':
		value := lxr.readString()
		token = Token{Type: TOKEN_STRING, Value: value}
	case ':':
		token = Token{Type: TOKEN_COLON, Value: ":"}
	case ',':
		token = Token{Type: TOKEN_COMMA, Value: ","}
	case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		value := lxr.readNumber()
		token = Token{Type: TOKEN_NUMBER, Value: value}
	case 't', 'f':
		tokenType, value := lxr.readBoolean()
		token = Token{Type: tokenType, Value: value}
	case 'n':
		value := lxr.readNull()
		token = Token{Type: TOKEN_NULL, Value: value}
	default:
		token.Type = TOKEN_EOF
	}

	lxr.nextChar()

	return token
}
