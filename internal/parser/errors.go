package parser

import (
	"fmt"
)

type ParseError struct {
	currentLine     int64
	currentPosition int64
	err             string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("[line: %d; pos: %d] parse error: %s", e.currentLine, e.currentPosition, e.err)
}

func (p *Parser) appendParseError(err string) {
	p.appendError(&ParseError{
		err:             err,
		currentLine:     p.lexer.CurrentLine(),
		currentPosition: p.lexer.CurrentPositionAtLine(),
	})
}

func (p *Parser) appendError(err error) {
	p.errors = append(p.errors, err)
}
