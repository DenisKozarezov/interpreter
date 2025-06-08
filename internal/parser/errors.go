package parser

import (
	"fmt"
)

type SyntaxError struct {
	err string
}

func wrapSyntaxError(err string) *SyntaxError {
	return &SyntaxError{err: err}
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("syntax error: %s", e.err)
}

type ParseError struct {
	err string
}

func wrapParseError(err string) *ParseError {
	return &ParseError{err: err}
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("parse error: %s", e.err)
}

func (p *Parser) appendSyntaxError(err string) {
	p.appendError(&SyntaxError{err: err})
}

func (p *Parser) appendParseError(err string) {
	p.appendError(&ParseError{err: err})
}

func (p *Parser) appendError(err error) {
	p.errors = append(p.errors, err)
}
