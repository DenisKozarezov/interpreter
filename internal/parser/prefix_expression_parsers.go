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

func (p *Parser) parseBooleanLiteral() ast.Expression {
	return &expressions.Boolean{Token: p.currentToken, Value: p.currentTokenIs(tokens.TRUE)}
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &expressions.PrefixExpression{Token: p.currentToken, Operator: p.currentToken.Literal}

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

func (p *Parser) parseBlockStatement() *expressions.BlockStatement {
	block := &expressions.BlockStatement{Token: p.currentToken}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for !p.currentTokenIs(tokens.RBRACE) && !p.currentTokenIs(tokens.EOF) {
		block.Statements = append(block.Statements, p.parseStatement())
		p.nextToken()
	}

	return block
}
