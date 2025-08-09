package tests

import (
	"interpreter/internal/ast/expressions"
	"interpreter/internal/ast/statements"
	"interpreter/internal/lexer/tokens"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArrayParsing(t *testing.T) {
	// 1. Arrange
	source := "[1, 2 * 2, 3 + 3]"

	// 2. Act
	statement := parseProgramAndCheck(t, source)

	// 3. Assert
	exp, ok := statement.(*statements.ExpressionStatement)
	require.Truef(t, ok, "expected an expression, got = %T", exp)
	array, ok := exp.Value.(*expressions.ArrayLiteral)
	require.Truef(t, ok, "expected an array, got = %T", array)
	require.Len(t, array.Items, 3, "expected 3 items in array")
	testIntegerExpression(t, array.Items[0], 1)
	testInfixExpression(t, array.Items[1], 2, tokens.ASTERISK, 2)
	testInfixExpression(t, array.Items[2], 3, tokens.PLUS, 3)
}
