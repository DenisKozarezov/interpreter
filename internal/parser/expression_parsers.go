package parser

import (
	"interpreter/internal/ast/statements"
	"interpreter/internal/lexer/tokens"
)

const (
	LOWEST      = iota + 1
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

type (
	prefixParserFn = func() statements.Expression
	infixParserFn  = func(expression statements.Expression) statements.Expression
)

func (p *Parser) initPrefixParsers() {
	p.prefixParseFns = map[tokens.TokenType]prefixParserFn{
		tokens.IDENTIFIER: p.parseIdentifier,
	}
}

func (p *Parser) initInfixParsers() {
	p.infixParseFns = map[tokens.TokenType]infixParserFn{}
}

func (p *Parser) parseExpression(precedence int) statements.Expression {
	prefix := p.prefixParseFns[p.currentToken.Type]
	if prefix == nil {
		return nil
	}
	leftExpression := prefix()

	return leftExpression
}

func (p *Parser) parseIdentifier() statements.Expression {
	return &statements.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
}
