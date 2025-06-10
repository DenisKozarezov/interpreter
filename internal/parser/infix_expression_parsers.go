package parser

import (
	"interpreter/internal/ast"
	"interpreter/internal/ast/expressions"
	"interpreter/internal/lexer/tokens"
)

type (
	infixParserFn = func(expression ast.Expression) ast.Expression
)

// Precedence означает порядок (ранг) оператора согласно арифметическим, семантическим
// правилам языка программирования. Чем выше ранг, тем выше порядок выполнения, а, следовательно
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

// Здесь представлены ранги от самого младшего (LOWEST) до самого старшего (CALL).
// LOWEST -> EQUALS -> LESSGREATER -> SUM -> PRODUCT -> PREFIX -> CALL
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
	tokens.EQ:       EQUALS,      // a == b;
	tokens.NOT_EQ:   EQUALS,      // a != b;
	tokens.LT:       LESSGREATER, // a < b;
	tokens.GT:       LESSGREATER, // a > b;
	tokens.PLUS:     SUM,         // a + b
	tokens.MINUS:    SUM,         // a - b;
	tokens.SLASH:    PRODUCT,     // a / b;
	tokens.ASTERISK: PRODUCT,     // a * b;
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

func (p *Parser) parseInfixExpression(leftExpression ast.Expression) ast.Expression {
	expression := &expressions.InfixExpression{
		Token:          p.currentToken,
		Operator:       p.currentToken.Literal,
		LeftExpression: leftExpression,
	}

	precedence := p.currentPrecedence()
	p.nextToken()
	expression.RightExpression = p.parseExpression(precedence)
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
