package parser

import (
	"fmt"

	"interpreter/internal/ast"
	"interpreter/internal/ast/statements"
	"interpreter/internal/lexer/tokens"
)

type lexer interface {
	NextToken() tokens.Token
}

type Parser struct {
	lexer  lexer
	errors []error

	currentToken tokens.Token
	peekToken    tokens.Token

	statementsParseFns map[tokens.TokenType]statementParserFn
	prefixParseFns     map[tokens.TokenType]prefixParserFn
	infixParseFns      map[tokens.TokenType]infixParserFn
}

func NewParser(lexer lexer) *Parser {
	parser := Parser{lexer: lexer}
	parser.initStatementParsers()
	parser.initPrefixParsers()
	parser.initInfixParsers()
	return &parser
}

func (p *Parser) Errors() []error {
	return p.errors
}

func (p *Parser) Parse() *ast.Program {
	program := ast.Program{Statements: make([]ast.Statement, 0)}

	p.nextToken()
	p.nextToken()

	for !p.currentTokenIs(tokens.EOF) {
		if statement := p.parseStatement(); statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}

	return &program
}

func (p *Parser) parseStatement() ast.Statement {
	statementFn, isStatement := p.statementsParseFns[p.currentToken.Type]
	if !isStatement {
		return p.parseExpressionStatement()
	}
	return statementFn()
}

func (p *Parser) parseExpressionStatement() ast.Statement {
	statement := statements.NewStatement(p.currentToken, p.parseExpression(LOWEST))

	if p.peekTokenIs(tokens.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

func (p *Parser) parseExpression(precedence Precedence) ast.Expression {
	prefix, prefixFound := p.prefixParseFns[p.currentToken.Type]
	if !prefixFound {
		p.appendParseError(fmt.Sprintf("no prefix parse function found for token `%s [%d]`", p.currentToken.Literal, p.currentToken.Type))
		return nil
	}
	expression := prefix()

	for !p.peekTokenIs(tokens.SEMICOLON) && precedence < p.peekPrecedence() {
		infix, infixFound := p.infixParseFns[p.peekToken.Type]
		if !infixFound {
			return expression
		}

		p.nextToken()

		expression = infix(expression)
	}

	return expression
}

func (p *Parser) expectToken(tokenType tokens.TokenType) bool {
	if p.peekTokenIs(tokenType) {
		p.nextToken()
		return true
	} else {
		p.appendParseError(fmt.Sprintf("expected token type [%d] but got [%d]", tokenType, p.peekToken.Type))
		return false
	}
}

func (p *Parser) peekTokenIs(tokenType tokens.TokenType) bool {
	return p.peekToken.Type == tokenType
}

func (p *Parser) currentTokenIs(tokenType tokens.TokenType) bool {
	return p.currentToken.Type == tokenType
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}
