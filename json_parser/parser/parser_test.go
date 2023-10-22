package parser_test

import (
	"os"
	"testing"

	"github.com/mauricioabreu/codingchallenges/json_parser/lexer"
	"github.com/mauricioabreu/codingchallenges/json_parser/parser"
	"github.com/stretchr/testify/assert"
)

func TestValidJSON(t *testing.T) {
	data, err := os.ReadFile("../tests/step1/valid.json")
	assert.NoError(t, err)
	psr := parser.NewParser(lexer.NewLexer(string(data)))
	assert.True(t, psr.Parse())
}

func TestInvalidJSON(t *testing.T) {
	data, err := os.ReadFile("../tests/step1/invalid.json")
	assert.NoError(t, err)
	psr := parser.NewParser(lexer.NewLexer(string(data)))
	assert.False(t, psr.Parse())
}
