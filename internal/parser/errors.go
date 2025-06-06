package parser

import (
	"fmt"

	"interpreter/internal/lexer/tokens"
)

func (p *Parser) appendTokenTypeError(tokenType tokens.TokenType) {
	p.appendError(fmt.Errorf("expected token type [%d] but got [%d]", tokenType, p.peekToken.Type))
}

func (p *Parser) appendError(err error) {
	p.errors = append(p.errors, err)
}
