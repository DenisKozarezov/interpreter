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
		tokens.TRUE:       p.parseBooleanLiteral,
		tokens.FALSE:      p.parseBooleanLiteral,
		tokens.LPAREN:     p.parseGroupedExpression,
		tokens.IF:         p.parseConditionExpression,
		tokens.FUNCTION:   p.parseFunction,
	}
}

func (p *Parser) parseIdentifier() ast.Expression {
	return expressions.NewIdentifier(p.currentToken)
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	literal := &expressions.IntegerLiteral{Token: p.currentToken}

	value, err := strconv.ParseInt(p.currentToken.Literal, 10, 64)
	if err != nil {
		p.parseError(fmt.Sprintf("could not parse '%q' as integer", p.currentToken.Literal))
		return nil
	}

	literal.Value = value
	return literal
}

func (p *Parser) parseBooleanLiteral() ast.Expression {
	return expressions.NewBoolean(p.currentToken)
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &expressions.PrefixExpression{Token: p.currentToken}

	p.nextToken()

	expression.RightExpression = p.parseExpression(PREFIX)
	return expression
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectToken(tokens.RPAREN) {
		return nil
	}

	return exp
}

func (p *Parser) parseConditionExpression() ast.Expression {
	expression := &expressions.ConditionExpression{Token: p.currentToken}

	if !p.expectToken(tokens.LPAREN) {
		return nil
	}

	expression.Condition = p.parseExpression(LOWEST)

	if !p.currentTokenIs(tokens.RPAREN) || !p.expectToken(tokens.LBRACE) {
		return nil
	}

	expression.Then = p.parseBlockStatement()

	if p.peekTokenIs(tokens.ELSE) {
		p.nextToken()

		if !p.expectToken(tokens.LBRACE) {
			return nil
		}

		expression.Else = p.parseBlockStatement()
	}

	return expression
}

func (p *Parser) parseFunction() ast.Expression {
	expression := &expressions.FunctionLiteral{Token: p.currentToken}

	if !p.expectToken(tokens.LPAREN) {
		return nil
	}

	expression.Args = p.parseFunctionArguments()

	if !p.expectToken(tokens.LBRACE) {
		return nil
	}

	expression.Body = p.parseBlockStatement()

	return expression
}

func (p *Parser) parseFunctionArguments() []*expressions.Identifier {
	var identifiers []*expressions.Identifier

	if p.peekTokenIs(tokens.RPAREN) {
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	identifiers = append(identifiers, expressions.NewIdentifier(p.currentToken))
	for p.peekTokenIs(tokens.COMMA) {
		p.nextToken()
		p.nextToken()
		identifiers = append(identifiers, expressions.NewIdentifier(p.currentToken))
	}

	if !p.expectToken(tokens.RPAREN) {
		return nil
	}

	return identifiers
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &expressions.CallExpression{Token: p.currentToken, Function: function}
	exp.Args = p.parseCallArguments()
	return exp
}

func (p *Parser) parseCallArguments() []ast.Expression {
	var args []ast.Expression

	if p.peekTokenIs(tokens.RPAREN) {
		p.nextToken()
		return args
	}

	p.nextToken()

	args = append(args, p.parseExpression(LOWEST))
	for p.peekTokenIs(tokens.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))
	}

	if !p.expectToken(tokens.RPAREN) {
		return nil
	}

	return args
}
