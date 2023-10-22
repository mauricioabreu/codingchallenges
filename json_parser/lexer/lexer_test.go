package lexer_test

import (
	"testing"

	"github.com/mauricioabreu/codingchallenges/json_parser/lexer"
	"github.com/stretchr/testify/assert"
)

func TestNextToken(t *testing.T) {
	lxr := lexer.NewLexer([]byte("{}"))
	assert.Equal(t, lxr.NextToken(), lexer.Token{Type: lexer.TOKEN_LEFT_BRACE, Value: "{"})
	assert.Equal(t, lxr.NextToken(), lexer.Token{Type: lexer.TOKEN_RIGHT_BRACE, Value: "}"})
}
