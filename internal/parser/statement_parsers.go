package parser

import (
	"interpreter/internal/ast/expressions"
	"interpreter/internal/ast/statements"
	"interpreter/internal/lexer/tokens"
)

type (
	statementParserFn = func() statements.Statement
)

func (p *Parser) initStatementParsers() {
	p.statementsParseFns = map[tokens.TokenType]statementParserFn{
		tokens.LET:    p.parseLetStatement,
		tokens.RETURN: p.parseReturnStatement,
	}
}

// parseLetStatement парсит конструкцию let.
func (p *Parser) parseLetStatement() statements.Statement {
	statement := statements.LetStatement{Token: p.currentToken}

	if !p.expectToken(tokens.IDENTIFIER) {
		return nil
	}

	statement.Identifier = expressions.NewIdentifier(p.currentToken)

	if !p.expectToken(tokens.ASSIGN) {
		return nil
	}
	p.nextToken()

	statement.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(tokens.SEMICOLON) {
		p.nextToken()
	}

	return &statement
}

// parseReturnStatement парсит конструкцию return.
func (p *Parser) parseReturnStatement() statements.Statement {
	statement := statements.ReturnStatement{Token: p.currentToken}

	p.nextToken()

	statement.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(tokens.SEMICOLON) {
		p.nextToken()
	}

	return &statement
}

func (p *Parser) parseBlockStatement() *statements.BlockStatement {
	block := &statements.BlockStatement{Token: p.currentToken}
	block.Statements = []statements.Statement{}

	p.nextToken()

	for !p.currentTokenIs(tokens.RBRACE) && !p.currentTokenIs(tokens.EOF) {
		statement := p.parseStatement()
		if statement != nil {
			block.Statements = append(block.Statements, statement)
		}
		p.nextToken()
	}

	return block
}
