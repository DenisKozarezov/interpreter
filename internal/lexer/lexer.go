package lexer

import (
	"bytes"
	"interpreter/internal/lexer/tokens"
	"io"
)

type Symbol = rune

const NULL Symbol = 0

type Reader interface {
	io.ReaderAt
}

type Lexer struct {
	reader Reader

	currentPosition   int64
	lineStartPosition int64
	currentLine       int64
	nextPosition      int64
	currentSymbol     Symbol
}

func NewLexer(reader Reader) *Lexer {
	l := &Lexer{reader: reader, currentSymbol: NULL, currentPosition: -1, nextPosition: 0, currentLine: 1}
	l.readSymbol()
	return l
}

func (l *Lexer) CurrentPositionAtLine() int64 {
	return l.currentPosition - l.lineStartPosition
}

func (l *Lexer) CurrentLine() int64 {
	return l.currentLine
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
	var buffer bytes.Buffer

	for fn(l.currentSymbol) {
		buffer.WriteRune(l.currentSymbol)
		l.readSymbol()
	}

	return buffer.String()
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

	if l.currentSymbol == '\n' {
		l.lineStartPosition = l.nextPosition
		l.currentLine++
	}
}

func (l *Lexer) peekSymbol() Symbol {
	bytes := make([]byte, 1)

	if _, err := l.reader.ReadAt(bytes, l.nextPosition); err == io.EOF {
		return NULL
	}

	return Symbol(bytes[0])
}
