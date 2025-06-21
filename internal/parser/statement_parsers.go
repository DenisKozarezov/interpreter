package parser

import (
	"interpreter/internal/ast"
	"interpreter/internal/ast/expressions"
	"interpreter/internal/ast/statements"
	"interpreter/internal/lexer/tokens"
)

type (
	statementParserFn = func() ast.Statement
)

func (p *Parser) initStatementParsers() {
	p.statementsParseFns = map[tokens.TokenType]statementParserFn{
		tokens.LET:    p.parseLetStatement,
		tokens.RETURN: p.parseReturnStatement,
	}
}

// parseLetStatement парсит конструкцию let.
func (p *Parser) parseLetStatement() ast.Statement {
	statement := statements.LetStatement{Token: p.currentToken}

	if !p.expectToken(tokens.IDENTIFIER) {
		return nil
	}

	statement.Identifier = expressions.NewIdentifier(p.currentToken)

	if !p.expectToken(tokens.ASSIGN) {
		return nil
	}

	p.skipUntilNextStatement()

	return &statement
}

// parseReturnStatement парсит конструкцию return.
func (p *Parser) parseReturnStatement() ast.Statement {
	statement := statements.ReturnStatement{Token: p.currentToken}

	p.nextToken()

	p.skipUntilNextStatement()

	return &statement
}
