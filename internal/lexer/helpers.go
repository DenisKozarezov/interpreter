package lexer

import (
	"unicode"
)

func isLetter(symbol Symbol) bool {
	return unicode.IsLetter(symbol)
}

func isDigit(symbol Symbol) bool {
	return unicode.IsDigit(symbol)
}

const (
	whitespace     = ' '
	tabulation     = '\t'
	newline        = '\n'
	carriageReturn = '\r'
	quot           = '"'
)

func isNewline(symbol Symbol) bool {
	return symbol == newline || symbol == carriageReturn
}
