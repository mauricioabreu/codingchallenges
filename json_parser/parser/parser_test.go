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

func TestStep2ValidJSON1(t *testing.T) {
	data, err := os.ReadFile("../tests/step2/valid.json")
	assert.NoError(t, err)
	psr := parser.NewParser(lexer.NewLexer(data))
	assert.True(t, psr.Parse())
}

func TestStep2ValidJSON2(t *testing.T) {
	data, err := os.ReadFile("../tests/step2/valid2.json")
	assert.NoError(t, err)
	psr := parser.NewParser(lexer.NewLexer(data))
	assert.True(t, psr.Parse())
}

func TestStep2InvalidJSON1(t *testing.T) {
	data, err := os.ReadFile("../tests/step2/invalid.json")
	assert.NoError(t, err)
	psr := parser.NewParser(lexer.NewLexer(data))
	assert.False(t, psr.Parse())
}

func TestStep2InvalidJSON2(t *testing.T) {
	data, err := os.ReadFile("../tests/step2/invalid2.json")
	assert.NoError(t, err)
	psr := parser.NewParser(lexer.NewLexer(data))
	assert.False(t, psr.Parse())
}

func TestStep3ValidJSON(t *testing.T) {
	data, err := os.ReadFile("../tests/step3/valid.json")
	assert.NoError(t, err)
	psr := parser.NewParser(lexer.NewLexer(data))
	assert.True(t, psr.Parse())
}

func TestStep3InvalidJSON(t *testing.T) {
	data, err := os.ReadFile("../tests/step3/invalid.json")
	assert.NoError(t, err)
	psr := parser.NewParser(lexer.NewLexer(data))
	assert.False(t, psr.Parse())
}

func TestStep4ValidJSON(t *testing.T) {
	data, err := os.ReadFile("../tests/step4/valid.json")
	assert.NoError(t, err)
	psr := parser.NewParser(lexer.NewLexer(data))
	assert.True(t, psr.Parse())
}
