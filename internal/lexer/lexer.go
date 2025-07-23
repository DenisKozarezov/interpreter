package lexer

import (
	"bytes"
	"io"

	"interpreter/internal/lexer/tokens"
)

type Symbol = rune

const NULL Symbol = 0

type Reader interface {
	io.ReaderAt
}

type Lexer struct {
	reader Reader

	currentSymbol     Symbol
	currentLine       int64
	currentPosition   int64
	nextPosition      int64
	lineStartPosition int64
}

func NewLexer(reader Reader) *Lexer {
	l := &Lexer{
		reader:          reader,
		currentLine:     1,
		currentSymbol:   NULL,
		currentPosition: -1,
		nextPosition:    0,
	}
	l.readSymbol()
	return l
}

func (l *Lexer) CurrentLine() int64 {
	return l.currentLine
}

func (l *Lexer) CurrentPositionAtLine() int64 {
	return l.currentPosition - l.lineStartPosition
}

func (l *Lexer) NextToken() tokens.Token {
	for {
		l.skipWhitespace()

		if l.currentSymbol == NULL {
			return tokens.NewToken(tokens.EOF, "")
		}

		literal := string(l.currentSymbol)
		currentTokenType, found := tokens.LookupTokenType(literal)
		if !found {
			return l.parseCustomToken(literal)
		}

		// Если текущий и следующий токены являются частью одного целого токена,
		// то объединяем их. Например:
		//  5 == 5 или 5 != 6
		//    ^^         ^^
		unitedLiteral := literal + string(l.peekSymbol())
		unitedTokenType, found := tokens.LookupTokenType(unitedLiteral)
		if found {
			literal = unitedLiteral
			currentTokenType = unitedTokenType

			switch unitedTokenType {
			case tokens.COMMENT_LINE:
				l.skipLine()
				continue
			case tokens.COMMENT_BEGIN:
				l.skipBlockComment()
				continue
			}

			l.readSymbol()
		}

		l.readSymbol()

		return tokens.NewToken(currentTokenType, literal)
	}
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
	for l.currentSymbol == whitespace || l.currentSymbol == tabulation || isNewline(l.currentSymbol) {
		l.readSymbol()
	}
}

func (l *Lexer) skipLine() {
	for !isNewline(l.currentSymbol) && l.currentSymbol != NULL {
		l.readSymbol()
	}
}

func (l *Lexer) skipBlockComment() {
	for l.peekSymbol() != NULL {
		if l.currentSymbol == '*' && l.peekSymbol() == '/' {
			l.readSymbol()
			l.readSymbol()
			break
		}
		l.readSymbol()
	}
}

func (l *Lexer) readSymbol() {
	l.currentSymbol = l.peekSymbol()
	l.currentPosition = l.nextPosition
	l.nextPosition++

	if l.currentSymbol == newline {
		l.currentLine++
		l.lineStartPosition = l.nextPosition
	}
}

func (l *Lexer) peekSymbol() Symbol {
	bytes := make([]byte, 1)

	if _, err := l.reader.ReadAt(bytes, l.nextPosition); err == io.EOF {
		return NULL
	}

	return Symbol(bytes[0])
}
