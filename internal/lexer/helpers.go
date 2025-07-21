package lexer

import (
	"unicode"

	"interpreter/internal/lexer/tokens"
)

func (l *Lexer) parseWord() (tokens.Token, bool) {
	if isLetter(l.currentSymbol) {
		literal := l.readLiteral(isLetter)
		tokenType := tokens.LookupIdentifierType(literal)
		return tokens.NewToken(tokenType, literal), true
	}
	return tokens.Token{}, false
}

func isLetter(symbol Symbol) bool {
	return unicode.IsLetter(symbol)
}

func (l *Lexer) parseDigit() (tokens.Token, bool) {
	if isDigit(l.currentSymbol) {
		return tokens.NewToken(tokens.INT, l.readLiteral(isDigit)), true
	}
	return tokens.Token{}, false
}

func isDigit(symbol Symbol) bool {
	return unicode.IsDigit(symbol)
}

const (
	whitespace     = ' '
	tabulation     = '\t'
	newline        = '\n'
	carriageReturn = '\r'
)

func isNewline(symbol Symbol) bool {
	return symbol == newline || symbol == carriageReturn
}
