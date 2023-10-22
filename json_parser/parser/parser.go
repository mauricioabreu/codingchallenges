package parser

import "github.com/mauricioabreu/codingchallenges/json_parser/lexer"

type Parser struct {
	lxr       *lexer.Lexer
	currToken lexer.Token
}

func NewParser(lxr *lexer.Lexer) *Parser {
	psr := &Parser{lxr: lxr}
	psr.nextToken()

	return psr
}

func (psr *Parser) nextToken() {
	psr.currToken = psr.lxr.NextToken()
}

func (psr *Parser) Parse() bool {
	if psr.currToken.Type != lexer.TOKEN_LEFT_BRACE {
		return false
	}

	psr.nextToken()

	if psr.currToken.Type != lexer.TOKEN_RIGHT_BRACE {
		return false
	}

	psr.nextToken()

	return true
}
