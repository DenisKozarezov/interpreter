package parser

import (
	"fmt"
)

type ParseError struct {
	currentPosAtLine int64
	err              string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("[pos: %d] parse error: %s", e.currentPosAtLine, e.err)
}

func (p *Parser) appendParseError(err string) {
	p.appendError(&ParseError{
		err:              err,
		currentPosAtLine: p.lexer.CurrentPositionAtLine(),
	})
}

func (p *Parser) appendError(err error) {
	p.errors = append(p.errors, err)
}
