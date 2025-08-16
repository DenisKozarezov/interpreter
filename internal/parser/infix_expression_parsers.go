package parser

import (
	"interpreter/internal/ast/expressions"
	"interpreter/internal/lexer/tokens"
)

type (
	infixParserFn = func(expression expressions.Expression) expressions.Expression
)

// Precedence means the order (rank) of an operator according to the arithmetic and
// semantic rules of the programming language. The higher the rank, the higher the
// execution order, and therefore the priority of the expression. For example, the
// expression (a + b + c) can be represented from left to right as shown below:
//
//	((a + b) + c)
//
// This is because both operators `+` have the same rank (SUM), and therefore they
// are executed sequentially. The expression (a + b * c) contains operators of different
// ranks, including the `*` (PRODUCT) operator. Therefore, we have:
//
//	(a + (b * c))
type Precedence = int8

// Ranks are presented here from the lowest (LOWEST) to the highest (INDEX).
// LOWEST -> EQUALS -> LESSGREATER -> SUM -> PRODUCT -> PREFIX -> CALL -> INDEX
const (
	LOWEST      Precedence = iota + 1
	EQUALS                 // ==
	LESSGREATER            // > or <
	SUM                    // +
	PRODUCT                // *
	PREFIX                 // -X or !X
	CALL                   // myFunction(X)
	INDEX                  // myArray[X]
)

var precedences = map[tokens.TokenType]Precedence{
	tokens.OR:       EQUALS,      // a || b
	tokens.EQ:       EQUALS,      // a == b;
	tokens.NOT_EQ:   EQUALS,      // a != b;
	tokens.LT:       LESSGREATER, // a < b;
	tokens.GT:       LESSGREATER, // a > b;
	tokens.LT_EQ:    LESSGREATER, // a <= b;
	tokens.GT_EQ:    LESSGREATER, // a >= b;
	tokens.AND:      LESSGREATER, // a && b
	tokens.PLUS:     SUM,         // a + b
	tokens.MINUS:    SUM,         // a - b;
	tokens.SLASH:    PRODUCT,     // a / b;
	tokens.ASTERISK: PRODUCT,     // a * b;
	tokens.LPAREN:   CALL,        // myFunction(arg1, arg2, ...)
	tokens.LBRACKET: INDEX,       // myArray[X]
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
		tokens.LT_EQ:    p.parseInfixExpression,
		tokens.GT_EQ:    p.parseInfixExpression,
		tokens.OR:       p.parseInfixExpression,
		tokens.AND:      p.parseInfixExpression,
		tokens.LPAREN:   p.parseCallExpression,
		tokens.LBRACKET: p.parseIndexExpression,
	}
}

func (p *Parser) parseInfixExpression(leftExpression expressions.Expression) expressions.Expression {
	expression := &expressions.InfixExpression{
		Token:          p.currentToken,
		LeftExpression: leftExpression,
	}

	precedence := p.currentPrecedence()
	p.nextToken()
	expression.RightExpression = p.parseExpression(precedence)
	return expression
}

func (p *Parser) parseIndexExpression(left expressions.Expression) expressions.Expression {
	expression := &expressions.IndexExpression{
		Token:          p.currentToken,
		LeftExpression: left,
	}

	p.nextToken()

	expression.Index = p.parseExpression(LOWEST)
	if !p.expectToken(tokens.RBRACKET) {
		return nil
	}

	return expression
}

func (p *Parser) currentPrecedence() Precedence {
	if precedence, found := precedences[p.currentToken.Type]; found {
		return precedence
	}
	return LOWEST
}

func (p *Parser) peekPrecedence() Precedence {
	if precedence, found := precedences[p.peekToken.Type]; found {
		return precedence
	}
	return LOWEST
}
