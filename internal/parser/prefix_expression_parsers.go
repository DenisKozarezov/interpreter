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

// parseIdentifier парсит некое строковое представление в идентификатор. Это может быть как
// идентификатор переменной, так и идентификатор функции и т.п. Например:
//
//	let x = 5;
//	    ^
//	let f = myFunc(x, y);
//	    ^    ^^^^  ^  ^
func (p *Parser) parseIdentifier() statements.Expression {
	return &statements.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
}

// parseIntegerLiteral парсит токен с типом INT в выражение, которое производит
// на свет некую целочисленную константу. Например:
//
//	5;
//	^
//	if 5 == 5 {
//	   ^    ^
//
// Важно понимать, что само по себе число 5 является лишь РЕЗУЛЬТАТОМ выражения, а
// не самим выражением. Это необходимо, чтобы были валидны следующие конструкции:
//
//	let y = 5;
//	let x = 5; let y = x;
//	let y = f(x);
//
// В последнем примере f(x) также является выражением, которое возвращает некое значение.
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
