package parser

import (
	"fmt"
	"strconv"

	"interpreter/internal/ast"
	"interpreter/internal/ast/expressions"
	"interpreter/internal/lexer/tokens"
)

type (
	prefixParserFn = func() ast.Expression
)

func (p *Parser) initPrefixParsers() {
	p.prefixParseFns = map[tokens.TokenType]prefixParserFn{
		tokens.IDENTIFIER: p.parseIdentifier,
		tokens.INT:        p.parseIntegerLiteral,
		tokens.BANG:       p.parsePrefixExpression,
		tokens.MINUS:      p.parsePrefixExpression,
	}
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &expressions.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	literal := &expressions.IntegerLiteral{Token: p.currentToken}

	value, err := strconv.ParseInt(p.currentToken.Literal, 10, 64)
	if err != nil {
		p.appendParseError(fmt.Sprintf("could not parse '%q' as integer", p.currentToken.Literal))
		return nil
	}

	literal.Value = value
	return literal
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &expressions.PrefixExpression{Token: p.currentToken, Operator: p.currentToken.Literal}

	p.nextToken()

	expression.RightExpression = p.parseExpression(PREFIX)
	return expression
}
