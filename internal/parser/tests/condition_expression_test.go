package tests

import (
	"interpreter/internal/ast/expressions"
	"interpreter/internal/ast/statements"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConditionExpression(t *testing.T) {
	// 1. Arrange
	source := `if (x < y) { x }`

	// 2. Act
	statement := parseProgramAndCheck(t, source)

	// 3. Assert
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
