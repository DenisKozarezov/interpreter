package lexer

import "interpreter/internal/lexer/tokens"

type Symbol = rune

const NULL Symbol = 0

type Lexer struct {
	source          string
	currentPosition int
	nextPosition    int
	currentSymbol   Symbol
}

func NewLexer(source string) *Lexer {
	lexer := &Lexer{source: source, currentSymbol: NULL}
	lexer.readSymbol()
	return lexer
}

func (l *Lexer) NextToken() tokens.Token {
	l.skipWhitespace()

	if l.currentSymbol == NULL {
		return tokens.NewToken(tokens.EOF, "")
	}

	sym := string(l.currentSymbol)
	tokenType, found := tokens.LookupTokenType(sym)
	if !found {
		return l.parseCustomToken(sym)
	}

	switch tokenType {
	case tokens.ASSIGN:
		if l.peekSymbol() == '=' {
			l.readSymbol()
			tokenType = tokens.EQ
			sym = "=="
		}
	case tokens.BANG:
		if l.peekSymbol() == '=' {
			l.readSymbol()
			tokenType = tokens.NOT_EQ
			sym = "!="
		}
	default:
	}

	l.readSymbol()

	return tokens.NewToken(tokenType, sym)
}

func (l *Lexer) parseCustomToken(literal string) tokens.Token {
	supportedCustomTokens := [...]func() (tokens.Token, bool){
		l.checkLetter,
		l.checkDigit,
	}

	for i := range supportedCustomTokens {
		if token, found := supportedCustomTokens[i](); found {
			return token
		}
	}

	return tokens.NewToken(tokens.ILLEGAL, literal)
}

func (l *Lexer) checkLetter() (tokens.Token, bool) {
	if isLetter(l.currentSymbol) {
		literal := l.readLiteral(isLetter)
		tokenType := tokens.LookupIdentifierType(literal)
		return tokens.NewToken(tokenType, literal), true
	}
	return tokens.Token{}, false
}

func (l *Lexer) checkDigit() (tokens.Token, bool) {
	if isDigit(l.currentSymbol) {
		return tokens.NewToken(tokens.INT, l.readLiteral(isDigit)), true
	}
	return tokens.Token{}, false
}

func (l *Lexer) readLiteral(fn func(Symbol) bool) string {
	startPosition := l.currentPosition
	for fn(l.currentSymbol) {
		l.readSymbol()
	}
	return l.source[startPosition:l.currentPosition]
}

func (l *Lexer) skipWhitespace() {
	for l.currentSymbol == ' ' || l.currentSymbol == '\t' || l.currentSymbol == '\n' || l.currentSymbol == '\r' {
		l.readSymbol()
	}
}

func (l *Lexer) readSymbol() {
	l.currentSymbol = l.peekSymbol()
	l.currentPosition = l.nextPosition
	l.nextPosition++
}

func (l *Lexer) peekSymbol() Symbol {
	if l.nextPosition >= len(l.source) {
		return NULL
	} else {
		return Symbol(l.source[l.nextPosition])
	}
}
