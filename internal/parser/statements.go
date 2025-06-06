package parser

import (
	"interpreter/internal/ast"
	"interpreter/internal/ast/statements"
	"interpreter/internal/lexer/tokens"
)

type statementParser = func() ast.Statement

func (p *Parser) initStatementParsers() {
	p.statementsParsers = map[tokens.TokenType]statementParser{
		tokens.LET:    p.parseLetStatement,
		tokens.RETURN: p.parseReturnStatement,
	}
}

func (p *Parser) parseLetStatement() ast.Statement {
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

func (p *Parser) parseReturnStatement() ast.Statement {
	statement := statements.ReturnStatement{Token: p.currentToken}

	p.nextToken()

	p.skipUntilNextStatement()

	return &statement
}
