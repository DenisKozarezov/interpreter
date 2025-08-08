package lexer

import (
	"strings"
	"testing"

	"interpreter/internal/lexer/tokens"

	"github.com/stretchr/testify/require"
)

func TestReadSymbol(t *testing.T) {
	for _, tt := range []struct {
		name           string
		source         string
		index          int64
		expectedSymbol Symbol
	}{
		{
			name:           "empty source",
			source:         "",
			index:          0,
			expectedSymbol: NULL,
		},
		{
			name:           "source with 1 symbol",
			source:         "A",
			index:          0,
			expectedSymbol: 'A',
		},
		{
			name:           "source with some symbols",
			source:         "ABCD",
			index:          1,
			expectedSymbol: 'B',
		},
		{
			name:           "last symbol",
			source:         "ABCD",
			index:          3,
			expectedSymbol: 'D',
		},
		{
			name:           "out of range",
			source:         "ABCD",
			index:          4,
			expectedSymbol: NULL,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			// 1. Arrange
			l := NewLexer(strings.NewReader(tt.source))

			// 2. Act
			for range tt.index {
				l.readSymbol()
			}

			// 3. Assert
			require.Equal(t, tt.expectedSymbol, l.currentSymbol, "current symbol should be equal")
			require.Equal(t, tt.index, l.currentPosition, "current position should equal")
			require.Equal(t, tt.index+1, l.nextPosition, "next position should equal")
		})
	}
}

func TestNextTokenWithWhitespaces(t *testing.T) {
	for _, tt := range []struct {
		name             string
		source           string
		expectedPosition int64
	}{
		{
			name:             "empty source",
			source:           "",
			expectedPosition: 0,
		},
		{
			name:             "1 tabulation",
			source:           "\t",
			expectedPosition: 1,
		},
		{
			name:             "2 tabulations",
			source:           "\t\t",
			expectedPosition: 2,
		},
		{
			name:             "newlines",
			source:           "\n\n",
			expectedPosition: 2,
		},
		{
			name:             "whitespaces",
			source:           "\n\n\t\t\rhello",
			expectedPosition: 10,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			// 1. Arrange
			l := NewLexer(strings.NewReader(tt.source))

			// 2. Act
			_ = l.NextToken()

			// 3. Assert
			require.Equal(t, tt.expectedPosition, l.currentPosition)
		})
	}
}

func TestNextToken(t *testing.T) {
	source := `let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
};

let result = add(five, ten);

!*-/5;
5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;
"foobar"
"foo bar"
`

	tests := []struct {
		expectedType    tokens.TokenType
		expectedLiteral string
	}{
		{tokens.LET, "let"},
		{tokens.IDENTIFIER, "five"},
		{tokens.ASSIGN, "="},
		{tokens.INT, "5"},
		{tokens.SEMICOLON, ";"},

		{tokens.LET, "let"},
		{tokens.IDENTIFIER, "ten"},
		{tokens.ASSIGN, "="},
		{tokens.INT, "10"},
		{tokens.SEMICOLON, ";"},

		{tokens.LET, "let"},
		{tokens.IDENTIFIER, "add"},
		{tokens.ASSIGN, "="},
		{tokens.FUNCTION, "fn"},
		{tokens.LPAREN, "("},
		{tokens.IDENTIFIER, "x"},
		{tokens.COMMA, ","},
		{tokens.IDENTIFIER, "y"},
		{tokens.RPAREN, ")"},
		{tokens.LBRACE, "{"},
		{tokens.IDENTIFIER, "x"},
		{tokens.PLUS, "+"},
		{tokens.IDENTIFIER, "y"},
		{tokens.SEMICOLON, ";"},
		{tokens.RBRACE, "}"},
		{tokens.SEMICOLON, ";"},

		{tokens.LET, "let"},
		{tokens.IDENTIFIER, "result"},
		{tokens.ASSIGN, "="},
		{tokens.IDENTIFIER, "add"},
		{tokens.LPAREN, "("},
		{tokens.IDENTIFIER, "five"},
		{tokens.COMMA, ","},
		{tokens.IDENTIFIER, "ten"},
		{tokens.RPAREN, ")"},
		{tokens.SEMICOLON, ";"},

		{tokens.BANG, "!"},
		{tokens.ASTERISK, "*"},
		{tokens.MINUS, "-"},
		{tokens.SLASH, "/"},
		{tokens.INT, "5"},
		{tokens.SEMICOLON, ";"},

		{tokens.INT, "5"},
		{tokens.LT, "<"},
		{tokens.INT, "10"},
		{tokens.GT, ">"},
		{tokens.INT, "5"},
		{tokens.SEMICOLON, ";"},

		{tokens.IF, "if"},
		{tokens.LPAREN, "("},
		{tokens.INT, "5"},
		{tokens.LT, "<"},
		{tokens.INT, "10"},
		{tokens.RPAREN, ")"},
		{tokens.LBRACE, "{"},
		{tokens.RETURN, "return"},
		{tokens.TRUE, "true"},
		{tokens.SEMICOLON, ";"},
		{tokens.RBRACE, "}"},
		{tokens.ELSE, "else"},
		{tokens.LBRACE, "{"},
		{tokens.RETURN, "return"},
		{tokens.FALSE, "false"},
		{tokens.SEMICOLON, ";"},
		{tokens.RBRACE, "}"},

		{tokens.INT, "10"},
		{tokens.EQ, "=="},
		{tokens.INT, "10"},
		{tokens.SEMICOLON, ";"},

		{tokens.INT, "10"},
		{tokens.NOT_EQ, "!="},
		{tokens.INT, "9"},
		{tokens.SEMICOLON, ";"},

		{tokens.STRING, "foobar"},
		{tokens.STRING, "foo bar"},

		{tokens.EOF, ""},
	}

	// 1. Arrange
	l := NewLexer(strings.NewReader(source))

	for _, tt := range tests {
		// 2. Act
		got := l.NextToken()

		// 3. Assert
		require.Equal(t, tt.expectedType, got.Type, "token type should be equal")
		require.Equal(t, tt.expectedLiteral, got.Literal, "token literal should be equal")
	}
}

func TestCommentLine(t *testing.T) {
	for _, tt := range []struct {
		name             string
		source           string
		expectedPosition int64
	}{
		{
			name:             "only commentary token appears in source, we start from zero index and skip 2 symbols",
			source:           `//`,
			expectedPosition: 2,
		},
		{
			name:             "whole line is commented, we skip all of it",
			source:           `// let x = 5;`,
			expectedPosition: 13,
		},
		{
			name: "first line is commented, we skip it and go straight to the let statement",
			source: `// comment
let x = 5;
`,
			expectedPosition: 14, // 10 symbols from the first line + 1 symbol of newline '\n' + 3 symbols of 'let'
		},
		{
			name: "many lines are commented, we skip all of them and go straight to the let statement",
			source: `// comment
// comment
// comment
// comment
let x = 5;
`,
			expectedPosition: 47, // 40 symbols from the commented lines + 4 symbols of newlines '\n' + 3 symbols of 'let'
		},
		{
			name:             "empty block comment",
			source:           `/**/`,
			expectedPosition: 4,
		},
		{
			name:             "one-line block comment",
			source:           `/* comment */`,
			expectedPosition: 13,
		},
		{
			name:             "one-line block comment 2",
			source:           `/*comment*/`,
			expectedPosition: 11,
		},
		{
			name: "multi-line block comment",
			source: `/*
comment
*/`,
			expectedPosition: 13,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			// 1. Arrange
			l := NewLexer(strings.NewReader(tt.source))

			// 2. Act
			_ = l.NextToken()

			// 3. Assert
			require.Equal(t, tt.expectedPosition, l.currentPosition)
		})
	}
}

func TestNewLine(t *testing.T) {
	for _, tt := range []struct {
		name         string
		source       string
		expectedLine int64
	}{
		{
			name:         "empty source - no lines",
			source:       ``,
			expectedLine: 1,
		},
		{
			name: "2 lines, but the first one is commented",
			source: `// some comment
let x = 2;`,
			expectedLine: 2,
		},
		{
			name: "many commented lines",
			source: `// some comment
// some comment
// some comment`,
			expectedLine: 3,
		},
		{
			name:         "first line + 3 newline symbols",
			source:       "\n\n\n",
			expectedLine: 4,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			// 1. Arrange
			l := NewLexer(strings.NewReader(tt.source))

			// 2. Act
			_ = l.NextToken()

			// 3. Assert
			require.Equal(t, tt.expectedLine, l.currentLine)
		})
	}
}
