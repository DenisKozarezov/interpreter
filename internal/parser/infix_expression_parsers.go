package parser

import (
	"interpreter/internal/ast/expressions"
	"interpreter/internal/lexer/tokens"
)

type (
	infixParserFn = func(expression expressions.Expression) expressions.Expression
)

// Precedence означает порядок (ранг) оператора согласно арифметическим, семантическим
// правилам языка программирования. Чем выше ранг, тем выше порядок выполнения, а, следовательно,
// приоритет выражения. Например, выражение (a + b + c) слева-направо можно представить
// как:
//
//	((a + b) + c)
//
// Потому что в выражении оба оператора `+` имеют равнозначные ранги (SUM), а следовательно выполняются
// последовательно. Выражение (a + b * c) содержит различные по рангу операторы, среди которых имеется
// оператор `*` (PRODUCT). Таким образом, получаем:
//
//	(a + (b * c))
type Precedence = int8

// Здесь представлены ранги от самого младшего (LOWEST) до самого старшего (INDEX).
// LOWEST -> EQUALS -> LESSGREATER -> SUM -> PRODUCT -> PREFIX -> CALL -> INDEX
const (
	LOWEST      Precedence = iota + 1
	EQUALS                 // ==
	LESSGREATER            // > or <
	SUM                    // +
	PRODUCT                // *
	PREFIX                 // -X or !X
	CALL                   // myFunction(X)
	INDEX
)

var precedences = map[tokens.TokenType]Precedence{
	tokens.EQ:       EQUALS,      // a == b;
	tokens.NOT_EQ:   EQUALS,      // a != b;
	tokens.LT:       LESSGREATER, // a < b;
	tokens.GT:       LESSGREATER, // a > b;
	tokens.PLUS:     SUM,         // a + b
	tokens.MINUS:    SUM,         // a - b;
	tokens.SLASH:    PRODUCT,     // a / b;
	tokens.ASTERISK: PRODUCT,     // a * b;
	tokens.LPAREN:   CALL,        // myFunction(arg1, arg2, ...)
	tokens.LBRACKET: INDEX,       // myArray[1]
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
