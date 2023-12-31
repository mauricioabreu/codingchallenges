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
	return psr.parseObject()
}

func (psr *Parser) parseObject() bool {
	if psr.currToken.Type != lexer.TOKEN_LEFT_BRACE {
		return false
	}

	psr.nextToken()

	for psr.currToken.Type == lexer.TOKEN_STRING {
		if !psr.parseKeyValue() {
			return false
		}

		// if next token is a comma, we have to look for another key-value pair
		if psr.currToken.Type == lexer.TOKEN_COMMA {
			psr.nextToken()

			if psr.currToken.Type != lexer.TOKEN_STRING {
				return false
			}
		}
	}

	if psr.currToken.Type != lexer.TOKEN_RIGHT_BRACE {
		return false
	}

	psr.nextToken()

	return true
}

func (psr *Parser) parseArray() bool {
	if psr.currToken.Type != lexer.TOKEN_LEFT_BRACKET {
		return false
	}

	psr.nextToken()

	// empty array
	if psr.currToken.Type == lexer.TOKEN_RIGHT_BRACKET {
		psr.nextToken()
		return true
	}

	for {
		if !psr.parseValue() {
			return false
		}

		if psr.currToken.Type == lexer.TOKEN_COMMA {
			psr.nextToken()
		} else if psr.currToken.Type == lexer.TOKEN_RIGHT_BRACKET {
			psr.nextToken()
			return true
		} else {
			return false
		}
	}
}

func (psr *Parser) parseKeyValue() bool {
	if psr.currToken.Type != lexer.TOKEN_STRING {
		return false
	}

	psr.nextToken()

	if psr.currToken.Type != lexer.TOKEN_COLON {
		return false
	}

	psr.nextToken()

	return psr.parseValue()
}

func (psr *Parser) parseValue() bool {
	switch psr.currToken.Type {
	case lexer.TOKEN_STRING, lexer.TOKEN_NUMBER, lexer.TOKEN_TRUE, lexer.TOKEN_FALSE, lexer.TOKEN_NULL:
		psr.nextToken()
		return true
	case lexer.TOKEN_LEFT_BRACE:
		return psr.parseObject()
	case lexer.TOKEN_LEFT_BRACKET:
		return psr.parseArray()
	default:
		return false
	}
}
