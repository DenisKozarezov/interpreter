package parser

import (
	"interpreter/internal/ast/statements"
	"interpreter/internal/lexer/tokens"
)

type (
	statementParserFn = func() statements.Statement
	prefixParserFn    = func() statements.Statement
	infixParserFn     = func(expression statements.Expression) statements.Statement
)

func (p *Parser) initStatementParsers() {
	p.statementsParsersFns = map[tokens.TokenType]statementParserFn{
		tokens.LET:    p.parseLetStatement,
		tokens.RETURN: p.parseReturnStatement,
	}
}

func (p *Parser) parseLetStatement() statements.Statement {
	statement := statements.LetStatement{Token: p.currentToken}

	if !p.expectToken(tokens.IDENTIFIER) {
		return nil
	}

	statement.Identifier = &statements.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	if !p.expectToken(tokens.ASSIGN) {
		return nil
	}

	p.skipUntilNextStatement()

	return &statement
}

func (p *Parser) parseReturnStatement() statements.Statement {
	statement := statements.ReturnStatement{Token: p.currentToken}

	p.nextToken()

	p.skipUntilNextStatement()

	return &statement
}
