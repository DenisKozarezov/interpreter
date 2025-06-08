package lexer

import "interpreter/internal/lexer/tokens"

func (l *Lexer) parseWord() (tokens.Token, bool) {
	if isLetter(l.currentSymbol) {
		literal := l.readLiteral(isLetter)
		tokenType := tokens.LookupIdentifierType(literal)
		return tokens.NewToken(tokenType, literal), true
	}
	return tokens.Token{}, false
}

func isLetter(symbol Symbol) bool {
	return 'a' <= symbol && symbol <= 'z' || 'A' <= symbol && symbol <= 'Z' || symbol == '_'
}

func (l *Lexer) parseDigit() (tokens.Token, bool) {
	if isDigit(l.currentSymbol) {
		return tokens.NewToken(tokens.INT, l.readLiteral(isDigit)), true
	}
	return tokens.Token{}, false
}

func isDigit(symbol Symbol) bool {
	return '0' <= symbol && symbol <= '9'
}
