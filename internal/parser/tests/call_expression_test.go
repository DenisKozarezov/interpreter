package tests

import (
	"interpreter/internal/ast/expressions"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCallExpression(t *testing.T) {
	// 1. Arrange
	source := `myFunc(1, 2 * 3, 4 + 5);`

	// 2. Act
	statement := parseProgramAndCheckExpression(t, source)

	// 3. Assert
	call, ok := statement.Value.(*expressions.CallExpression)
	require.True(t, ok, "expected call expression")
	testIdentifier(t, call.Function, "myFunc")
	require.Len(t, call.Args, 3, "expected 3 arguments passed in function")
	testLiteralExpression(t, call.Args[0], 1)
	testInfixExpression(t, call.Args[1], 2, "*", 3)
	testInfixExpression(t, call.Args[2], 4, "+", 5)
}
