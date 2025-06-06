package parser

import (
	"interpreter/internal/ast"
	"interpreter/internal/lexer/tokens"
)

type lexer interface {
	NextToken() tokens.Token
}

type Parser struct {
	lexer             lexer
	errors            []error
	statementsParsers map[tokens.TokenType]statementParser

	currentToken tokens.Token
	peekToken    tokens.Token
}

func NewParser(lexer lexer) *Parser {
	parser := Parser{lexer: lexer}
	parser.initStatementParsers()
	parser.nextToken()
	parser.nextToken()
	return &parser
}

func (p *Parser) Errors() []error {
	return p.errors
}

func (p *Parser) Parse() *ast.Program {
	program := ast.Program{Statements: make([]ast.Statement, 0)}

	for !p.currentTokenIs(tokens.EOF) {
		if statement := p.parseStatement(); statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}

	return &program
}

func (p *Parser) parseStatement() ast.Statement {
	statement, found := p.statementsParsers[p.currentToken.Type]
	if !found {
		return nil
	}
	return statement()
}

func (p *Parser) expectToken(tokenType tokens.TokenType) bool {
	if p.peekTokenIs(tokenType) {
		p.nextToken()
		return true
	} else {
		p.appendTokenTypeError(tokenType)
		return false
	}
}

func (p *Parser) peekTokenIs(tokenType tokens.TokenType) bool {
	return p.peekToken.Type == tokenType
}

func (p *Parser) skipUntilNextStatement() {
	for !p.currentTokenIs(tokens.SEMICOLON) {
		p.nextToken()
	}
}

func (p *Parser) currentTokenIs(tokenType tokens.TokenType) bool {
	return p.currentToken.Type == tokenType
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}
