package tests

import (
	"fmt"
	"strconv"
	"testing"

	"interpreter/internal/ast"
	"interpreter/internal/ast/expressions"

	"github.com/stretchr/testify/require"
)

func TestIdentifierExpression(t *testing.T) {
	// 1. Arrange
	source := `foobar;`

	// 2. Act
	statement := parseProgramAndCheckExpression(t, source)

	// 3. Assert
	testLiteralExpression(t, statement.Value, "foobar")
}

func TestIntegerLiteralExpression(t *testing.T) {
	// 1. Arrange
	source := `5;`

	// 2. Act
	statement := parseProgramAndCheckExpression(t, source)

	// 3. Assert
	testLiteralExpression(t, statement.Value, 5)
}

func TestPrefixExpression(t *testing.T) {
	for _, tt := range []struct {
		source           string
		expectedOperator string
		rightExpression  any
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
		{"!true;", "!", true},
		{"!false;", "!", false},
	} {
		t.Run(tt.source, func(t *testing.T) {
			// 1. Act
			statement := parseProgramAndCheckExpression(t, tt.source)

			// 2. Assert
			prefix, ok := statement.Value.(*expressions.PrefixExpression)
			require.True(t, ok, "expression is not a prefix")
			require.Equal(t, tt.expectedOperator, prefix.Operator)

			testLiteralExpression(t, prefix.RightExpression, tt.rightExpression)
		})
	}
}

func TestInfixExpression(t *testing.T) {
	for _, tt := range []struct {
		source           string
		leftExpression   any
		expectedOperator string
		rightExpression  any
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
		{"true == true", true, "==", true},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
	} {
		t.Run(tt.source, func(t *testing.T) {
			// 1. Act
			statement := parseProgramAndCheckExpression(t, tt.source)

			// 2. Assert
			testInfixExpression(t, statement.Value, tt.leftExpression, tt.expectedOperator, tt.rightExpression)
		})
	}
}

func TestBooleanExpression(t *testing.T) {
	for _, tt := range []struct {
		source   string
		expected bool
	}{
		{"true;", true},
		{"false;", false},
	} {
		t.Run(tt.source, func(t *testing.T) {
			// 1. Act
			statement := parseProgramAndCheckExpression(t, tt.source)

			// 2. Assert
			testLiteralExpression(t, statement.Value, tt.expected)
		})
	}
}

func testInfixExpression(t *testing.T, exp ast.Expression, left any, op string, right any) {
	infix, ok := exp.(*expressions.InfixExpression)
	require.True(t, ok, "expression is not an infix")
	require.Equal(t, op, infix.Operator)

	testLiteralExpression(t, infix.LeftExpression, left)
	testLiteralExpression(t, infix.RightExpression, right)
}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected any) {
	switch v := expected.(type) {
	case int:
		testIntegerExpression(t, exp, int64(v))
		return
	case int64:
		testIntegerExpression(t, exp, v)
		return
	case string:
		testIdentifier(t, exp, v)
		return
	case bool:
		testBooleanLiteral(t, exp, v)
		return
	default:
		t.Fatalf("unexpected type")
	}
}

func testIdentifier(t *testing.T, exp ast.Expression, expected string) {
	ident, ok := exp.(*expressions.Identifier)
	require.True(t, ok, "expression is not an identifier")
	require.Equal(t, expected, ident.Literal())
}

func testIntegerExpression(t *testing.T, exp ast.Expression, expected int64) {
	integer, ok := exp.(*expressions.IntegerLiteral)
	require.True(t, ok, "expression is not an integer")
	require.Equal(t, expected, integer.Value)
	require.Equal(t, strconv.FormatInt(expected, 10), integer.Literal())
}

func testBooleanLiteral(t *testing.T, exp ast.Expression, expected bool) {
	boolean, ok := exp.(*expressions.Boolean)
	require.True(t, ok, "expression is not a boolean")
	require.Equal(t, expected, boolean.Value)
	require.Equal(t, fmt.Sprintf("%t", expected), boolean.Literal())
}
