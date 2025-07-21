package tests

import (
	"testing"

	"interpreter/internal/ast/expressions"
	"interpreter/internal/ast/statements"
	"interpreter/internal/lexer/tokens"

	"github.com/stretchr/testify/require"
)

func TestFunctionLiteral(t *testing.T) {
	// 1. Arrange
	source := `fn(x, y) { x + y }`

	// 2. Act
	statement := parseProgramAndCheckExpression(t, source)

	// 3. Assert
	fn, ok := statement.Value.(*expressions.FunctionLiteral)
	require.True(t, ok, "expected function literal")
	require.Len(t, fn.Args, 2, "expected 2 arguments")
	require.Equal(t, fn.Args[0].Literal(), "x")
	require.Equal(t, fn.Args[1].Literal(), "y")
	require.NotNil(t, fn.Body)
	body, ok := fn.Body.(*statements.BlockStatement)
	require.True(t, ok, "expected block statement")
	require.Len(t, body.Statements, 1, "expected some statements in function's body")

	exp, ok := body.Statements[0].(*statements.ExpressionStatement)
	require.True(t, ok, "expected expression in function's body")
	testInfixExpression(t, exp.Value, "x", tokens.PLUS, "y")
}

func TestFunctionArguments(t *testing.T) {
	for _, tt := range []struct {
		source       string
		expectedArgs []string
	}{
		{source: "fn() {};", expectedArgs: []string{}},
		{source: "fn(x) {};", expectedArgs: []string{"x"}},
		{source: "fn(x, y) {};", expectedArgs: []string{"x", "y"}},
		{source: "fn(x, y, z) {};", expectedArgs: []string{"x", "y", "z"}},
	} {
		t.Run(tt.source, func(t *testing.T) {
			// 1. Act
			statement := parseProgramAndCheckExpression(t, tt.source)

			// 2. Assert
			fn, ok := statement.Value.(*expressions.FunctionLiteral)
			require.True(t, ok, "expected function literal")
			require.Len(t, fn.Args, len(tt.expectedArgs), "function must have same arguments as expected")

			for i := range tt.expectedArgs {
				testLiteralExpression(t, fn.Args[i], tt.expectedArgs[i])
			}
		})
	}
}
