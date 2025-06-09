package parser

import (
	"fmt"
)

type SyntaxError struct {
	err string
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("syntax error: %s", e.err)
}

type ParseError struct {
	err string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("parse error: %s", e.err)
}

func (p *Parser) appendParseError(err string) {
	p.appendError(&ParseError{err: err})
}

func (p *Parser) appendError(err error) {
	p.errors = append(p.errors, err)
}
