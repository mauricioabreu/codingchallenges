package parser_test

import (
	"os"
	"testing"

	"github.com/mauricioabreu/codingchallenges/json_parser/lexer"
	"github.com/mauricioabreu/codingchallenges/json_parser/parser"
	"github.com/stretchr/testify/assert"
)

func TestStep1ValidJSON(t *testing.T) {
	data, err := os.ReadFile("../tests/step1/valid.json")
	assert.NoError(t, err)
	psr := parser.NewParser(lexer.NewLexer(data))
	assert.True(t, psr.Parse())
}

func TestStep1InvalidJSON(t *testing.T) {
	data, err := os.ReadFile("../tests/step1/invalid.json")
	assert.NoError(t, err)
	psr := parser.NewParser(lexer.NewLexer(data))
	assert.False(t, psr.Parse())
}

func TestStep2ValidJSON(t *testing.T) {
	data, err := os.ReadFile("../tests/step2/valid.json")
	assert.NoError(t, err)
	psr := parser.NewParser(lexer.NewLexer(data))
	assert.True(t, psr.Parse())
}
