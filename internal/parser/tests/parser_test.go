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

	ident, ok := statement.Value.(*statements.IntegerLiteral)
	require.True(t, ok, "expression is not an integer")
	require.Equal(t, "5", ident.Literal())
}
