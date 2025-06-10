package parser

import (
	"fmt"
	"strconv"

	"interpreter/internal/ast/statements"
	"interpreter/internal/lexer/tokens"
)

type (
	prefixParserFn = func() statements.Expression
)

func (p *Parser) initPrefixParsers() {
	p.prefixParseFns = map[tokens.TokenType]prefixParserFn{
		tokens.IDENTIFIER: p.parseIdentifier,
		tokens.INT:        p.parseIntegerLiteral,
		tokens.BANG:       p.parsePrefixExpression,
		tokens.MINUS:      p.parsePrefixExpression,
	}
}

func (p *Parser) parseIdentifier() statements.Expression {
	return &statements.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
}

func (p *Parser) parseIntegerLiteral() statements.Expression {
	literal := &statements.IntegerLiteral{Token: p.currentToken}

	value, err := strconv.ParseInt(p.currentToken.Literal, 10, 64)
	if err != nil {
		p.appendParseError(fmt.Sprintf("could not parse '%q' as integer", p.currentToken.Literal))
		return nil
	}

	literal.Value = value
	return literal
}

func (p *Parser) parsePrefixExpression() statements.Expression {
	expression := &statements.PrefixExpression{Token: p.currentToken, Operator: p.currentToken.Literal}

	p.nextToken()

	expression.RightExpression = p.parseExpression(PREFIX)
	return expression
}
