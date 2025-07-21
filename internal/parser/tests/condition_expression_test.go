package tests

import (
	"testing"

	"interpreter/internal/ast/expressions"
	"interpreter/internal/ast/statements"
	"interpreter/internal/lexer/tokens"

	"github.com/stretchr/testify/require"
)

func TestConditionExpression(t *testing.T) {
	// 1. Arrange
	source := `if (x < y) { x }`

	// 2. Act
	statement := parseProgramAndCheckExpression(t, source)

	// 3. Assert
	condition, ok := statement.Value.(*expressions.ConditionExpression)
	require.True(t, ok, "expected `if` statement")
	require.NotNil(t, condition.Then, "`then` block must have some body")
	then, ok := condition.Then.(*statements.BlockStatement)
	require.True(t, ok, "expected block statement")
	require.Len(t, then.Statements, 1)
	require.Nil(t, condition.Else, "`else` block must be empty")
	testInfixExpression(t, condition.Condition, "x", tokens.LT, "y")

	statement, ok = then.Statements[0].(*statements.ExpressionStatement)
	require.True(t, ok)
	testIdentifier(t, statement.Value, "x")
}
