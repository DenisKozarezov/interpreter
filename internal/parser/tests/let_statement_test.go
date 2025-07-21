package tests

import (
	"strings"
	"testing"

	"interpreter/internal/lexer"
	"interpreter/internal/parser"

	"interpreter/internal/ast/statements"

	"github.com/stretchr/testify/require"
)

func TestLetStatement(t *testing.T) {
	for _, tt := range []struct {
		source             string
		expectedIdentifier string
		expectedValue      any
	}{
		{"let x = 5;", "x", 5},
		{"let y = true;", "y", true},
		{"let foobar = y;", "foobar", "y"},
	} {
		t.Run(tt.source, func(t *testing.T) {
			// 1. Act
			statement := parseProgramAndCheck(t, tt.source)

			// 2. Assert
			let, ok := statement.(*statements.LetStatement)
			require.True(t, ok, "expected let statement")
			require.Equal(t, tt.expectedIdentifier, let.Identifier.Literal(), "expected identifier literal")
			testLiteralExpression(t, let.Value, tt.expectedValue)
		})
	}
}

func parseProgramAndCheckExpression(t *testing.T, source string) *statements.ExpressionStatement {
	// 1. Act
	statement := parseProgramAndCheck(t, source)

	// 2. Assert
	expression, ok := statement.(*statements.ExpressionStatement)
	require.True(t, ok, "statement is not an expression")

	return expression
}

func parseProgramAndCheck(t *testing.T, source string) statements.Statement {
	// 1. Arrange
	l := lexer.NewLexer(strings.NewReader(source))
	p := parser.NewParser(l)

	// 2. Act
	program := p.Parse()

	// 3. Assert
	require.Len(t, p.Errors(), 0)
	require.Len(t, program.Statements, 1)

	return program.Statements[0]
}
