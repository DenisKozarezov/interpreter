package tests

import (
	"interpreter/internal/ast/expressions"
	"interpreter/internal/ast/statements"
	"interpreter/internal/lexer"
	"interpreter/internal/parser"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConditionExpression(t *testing.T) {
	source := `if (x < y) { x }`

	// 1. Arrange
	l := lexer.NewLexer(strings.NewReader(source))
	p := parser.NewParser(l)

	// 2. Act
	program := p.Parse()

	// 3. Assert
	require.Len(t, p.Errors(), 0)
	require.Len(t, program.Statements, 1)

	statement, ok := program.Statements[0].(*statements.ExpressionStatement)
	require.True(t, ok, "expected expression statement")

	condition, ok := statement.Value.(*expressions.ConditionExpression)
	require.True(t, ok, "expected `if` statement")
	require.NotNil(t, condition.Then, "`then` block must have some body")
	require.Len(t, condition.Then.Statements, 1)
	require.Nil(t, condition.Else, "`else` block must be empty")
	testInfixExpression(t, condition.Condition, "x", "<", "y")

	statement, ok = condition.Then.Statements[0].(*statements.ExpressionStatement)
	require.True(t, ok)
	testIdentifier(t, statement.Value, "x")
}
