package tests

import (
	"github.com/stretchr/testify/require"
	"interpreter/internal/ast/statements"
	"interpreter/internal/lexer"
	"interpreter/internal/parser"
	"testing"
)

func TestIdentifierExpression(t *testing.T) {
	source := `foobar;`

	l := lexer.NewLexer(source)
	p := parser.NewParser(l)
	program := p.Parse()

	require.Len(t, program.Statements, 1)

	statement, ok := program.Statements[0].(*statements.ExpressionStatement)
	require.True(t, ok, "statement is not an expression")

	ident, ok := statement.Value.(*statements.Identifier)
	require.True(t, ok, "statement is not an identifier")
	require.Equal(t, "foobar", ident.Literal())
}
