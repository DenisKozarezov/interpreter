package tests

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"interpreter/internal/ast/statements"
	"interpreter/internal/lexer"
	"interpreter/internal/parser"
)

func TestIdentifierExpression(t *testing.T) {
	// 1. Arrange
	source := `foobar;`

	l := lexer.NewLexer(strings.NewReader(source))
	p := parser.NewParser(l)

	// 2. Act
	program := p.Parse()

	// 3. Assert
	require.Len(t, p.Errors(), 0)
	require.Len(t, program.Statements, 1)

	statement, ok := program.Statements[0].(*statements.ExpressionStatement)
	require.True(t, ok, "statement is not an expression")

	ident, ok := statement.Value.(*statements.Identifier)
	require.True(t, ok, "expression is not an identifier")
	require.Equal(t, "foobar", ident.Literal())
}

func TestIntegerLiteralExpression(t *testing.T) {
	// 1. Arrange
	source := `5;`

	l := lexer.NewLexer(strings.NewReader(source))
	p := parser.NewParser(l)

	// 2. Act
	program := p.Parse()

	// 3. Assert
	require.Len(t, p.Errors(), 0)
	require.Len(t, program.Statements, 1)

	statement, ok := program.Statements[0].(*statements.ExpressionStatement)
	require.True(t, ok, "statement is not an expression")

	checkIntegerExpression(t, statement.Value, 5)
}

func TestPrefixExpression(t *testing.T) {
	for _, tt := range []struct {
		name             string
		source           string
		expectedOperator string
		expectedValue    int64
	}{
		{
			name:             "number",
			source:           "!5;",
			expectedOperator: "!",
			expectedValue:    5,
		},
		{
			name:             "minus",
			source:           "-15;",
			expectedOperator: "-",
			expectedValue:    15,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			// 1. Arrange
			l := lexer.NewLexer(strings.NewReader(tt.source))
			p := parser.NewParser(l)

			// 2. Act
			program := p.Parse()

			// 3. Assert
			require.Len(t, p.Errors(), 0)
			require.Len(t, program.Statements, 1)

			statement, ok := program.Statements[0].(*statements.ExpressionStatement)
			require.True(t, ok, "statement is not an expression")

			prefix, ok := statement.Value.(*statements.PrefixExpression)
			require.True(t, ok, "expression is not a prefix")
			require.Equal(t, tt.expectedOperator, prefix.Operator)

			checkIntegerExpression(t, prefix.RightExpression, tt.expectedValue)
		})
	}
}

func checkIntegerExpression(t *testing.T, exp statements.Expression, value int64) {
	integer, ok := exp.(*statements.IntegerLiteral)
	require.True(t, ok, "expression is not an integer")
	require.Equal(t, value, integer.Value)
}
