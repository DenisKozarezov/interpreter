package parser

import (
	"interpreter/internal/ast/statements"
	"interpreter/internal/lexer/tokens"
)

type (
	infixParserFn = func(expression statements.Expression) statements.Expression
)

type Precedence = int8

const (
	LOWEST      Precedence = iota + 1
	EQUALS                 // ==
	LESSGREATER            // > or <
	SUM                    // +
	PRODUCT                // *
	PREFIX                 // -X or !X
	CALL                   // myFunction(X)
)

var precedences = map[tokens.TokenType]Precedence{
	tokens.EQ:       EQUALS,
	tokens.NOT_EQ:   EQUALS,
	tokens.LT:       LESSGREATER,
	tokens.GT:       LESSGREATER,
	tokens.PLUS:     SUM,
	tokens.MINUS:    SUM,
	tokens.SLASH:    PRODUCT,
	tokens.ASTERISK: PRODUCT,
}

func (p *Parser) initInfixParsers() {
	p.infixParseFns = map[tokens.TokenType]infixParserFn{
		tokens.PLUS:     p.parseInfixExpression,
		tokens.MINUS:    p.parseInfixExpression,
		tokens.SLASH:    p.parseInfixExpression,
		tokens.ASTERISK: p.parseInfixExpression,
		tokens.EQ:       p.parseInfixExpression,
		tokens.NOT_EQ:   p.parseInfixExpression,
		tokens.LT:       p.parseInfixExpression,
		tokens.GT:       p.parseInfixExpression,
	}
}

func (p *Parser) peekPrecedence() Precedence {
	if precedence, found := precedences[p.peekToken.Type]; found {
		return precedence
	}
	return LOWEST
}

func (p *Parser) currentPrecedence() Precedence {
	if precedence, found := precedences[p.currentToken.Type]; found {
		return precedence
	}
	return LOWEST
}

func (p *Parser) parseInfixExpression(leftExpression statements.Expression) statements.Expression {
	expression := &statements.InfixExpression{
		Token:          p.currentToken,
		Operator:       p.currentToken.Literal,
		LeftExpression: leftExpression,
	}

	precedence := p.currentPrecedence()
	p.nextToken()
	expression.RightExpression = p.parseExpression(precedence)
	return expression
}
