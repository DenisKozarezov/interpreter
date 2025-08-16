package parser

import (
	"fmt"
)

type ParseError struct {
	currentLine      int16
	currentPosAtLine int64
	err              string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("[line: %d; pos: %d] parse error: %s", e.currentLine, e.currentPosAtLine, e.err)
}

func (p *Parser) parseError(err string) {
	p.appendError(&ParseError{
		err:              err,
		currentLine:      p.lexer.CurrentLine(),
		currentPosAtLine: p.lexer.CurrentPositionAtLine(),
	})
}

func (p *Parser) appendError(err error) {
	p.errors = append(p.errors, err)
}
