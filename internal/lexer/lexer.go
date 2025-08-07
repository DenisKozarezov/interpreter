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
	var token tokens.Token

	for {
		l.skipWhitespace()

		currentSym := string(l.currentSymbol)

		switch l.currentSymbol {
		case '=':
			if l.peekSymbol() == '=' {
				ch := l.currentSymbol
				l.readSymbol()
				literal := string(ch) + string(l.currentSymbol)
				token = tokens.NewToken(tokens.EQ, literal)
			} else {
				token = tokens.NewToken(tokens.ASSIGN, currentSym)
			}
		case '+':
			token = tokens.NewToken(tokens.PLUS, currentSym)
		case '-':
			token = tokens.NewToken(tokens.MINUS, currentSym)
		case '!':
			if l.peekSymbol() == '=' {
				ch := l.currentSymbol
				l.readSymbol()
				literal := string(ch) + string(l.currentSymbol)
				token = tokens.NewToken(tokens.NOT_EQ, literal)
			} else {
				token = tokens.NewToken(tokens.BANG, currentSym)
			}
		case '/':
			if l.peekSymbol() == '/' {
				l.skipLine()
				continue
			} else if l.peekSymbol() == '*' {
				l.skipBlockComment()
				continue
			} else {
				token = tokens.NewToken(tokens.SLASH, currentSym)
			}
		case '*':
			token = tokens.NewToken(tokens.ASTERISK, currentSym)
		case '<':
			token = tokens.NewToken(tokens.LT, currentSym)
		case '>':
			token = tokens.NewToken(tokens.GT, currentSym)
		case ';':
			token = tokens.NewToken(tokens.SEMICOLON, currentSym)
		case ',':
			token = tokens.NewToken(tokens.COMMA, currentSym)
		case '{':
			token = tokens.NewToken(tokens.LBRACE, currentSym)
		case '}':
			token = tokens.NewToken(tokens.RBRACE, currentSym)
		case '(':
			token = tokens.NewToken(tokens.LPAREN, currentSym)
		case ')':
			token = tokens.NewToken(tokens.RPAREN, currentSym)
		case '[':
			token = tokens.NewToken(tokens.LBRACKET, currentSym)
		case ']':
			token = tokens.NewToken(tokens.RBRACKET, currentSym)
		case quot:
			l.readSymbol()
			token = tokens.NewToken(tokens.STRING, l.readRawString())
		case NULL:
			return tokens.NewToken(tokens.EOF, "")
		default:
			if isLetter(l.currentSymbol) {
				literal := l.readLiteral(isLetter)
				return tokens.NewToken(tokens.LookupIdentifierType(literal), literal)
			} else if isDigit(l.currentSymbol) {
				return tokens.NewToken(tokens.INT, l.readLiteral(isDigit))
			} else {
				token = tokens.NewToken(tokens.ILLEGAL, currentSym)
			}
		}

		l.readSymbol()
		return token
	}
}

func (l *Lexer) readLiteral(fn func(Symbol) bool) string {
	var buffer bytes.Buffer

	for l.currentSymbol != NULL && fn(l.currentSymbol) {
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

func (l *Lexer) readRawString() string {
	return l.readLiteral(func(symbol Symbol) bool {
		return symbol != quot
	})
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
