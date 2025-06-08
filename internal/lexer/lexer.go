package lexer

import (
	"interpreter/internal/lexer/tokens"
)

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

	currentSym := string(l.currentSymbol)
	currentTokenType, found := tokens.LookupTokenType(currentSym)
	if !found {
		return l.parseCustomToken(currentSym)
	}

	// Если текущий и следующий токены являются частью одного целого токена,
	// то объединяем их. Например:
	//  5 == 5 или 5 != 6
	//    ^^         ^^
	unitedSym := currentSym + string(l.peekSymbol())
	unitedTokenType, found := tokens.LookupTokenType(unitedSym)
	if found {
		currentSym = unitedSym
		currentTokenType = unitedTokenType
		l.readSymbol()
	}

	l.readSymbol()

	return tokens.NewToken(currentTokenType, currentSym)
}

func (l *Lexer) parseCustomToken(literal string) tokens.Token {
	supportedCustomTokensParseFns := [...]func() (tokens.Token, bool){
		l.parseWord,
		l.parseDigit,
	}

	for i := range supportedCustomTokensParseFns {
		if token, found := supportedCustomTokensParseFns[i](); found {
			return token
		}
	}

	return tokens.NewToken(tokens.ILLEGAL, literal)
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
