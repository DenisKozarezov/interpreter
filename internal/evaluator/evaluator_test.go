package evaluator

import (
	"strings"
	"testing"

	"interpreter/internal/lexer"
	"interpreter/internal/object"
	"interpreter/internal/parser"

	"github.com/stretchr/testify/require"
)

func TestEvalIntegerExpression(t *testing.T) {
	for _, tt := range []struct {
		source   string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	} {
		t.Run(tt.source, func(t *testing.T) {
			// 1. Act
			got := testEval(t, tt.source)

			// 2. Assert
			testIntegerObject(t, got, tt.expected)
		})
	}
}

func testEval(t *testing.T, source string) object.Object {
	// 1. Arrange
	l := lexer.NewLexer(strings.NewReader(source))
	p := parser.NewParser(l)

	// 2. Act
	program := p.Parse()

	// 3. Assert
	require.Len(t, p.Errors(), 0)
	require.Greater(t, len(program.Statements), 0)

	return EvaluateStatement(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) {
	result, ok := obj.(*object.Integer)
	require.True(t, ok, "expected integer object")
	require.Equal(t, expected, result.Value, "object's value is wrong")
}

func TestEvalBoolean(t *testing.T) {
	for _, tt := range []struct {
		source   string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
	} {
		t.Run(tt.source, func(t *testing.T) {
			// 1. Act
			got := testEval(t, tt.source)

			// 2. Assert
			testBooleanObject(t, got, tt.expected)
		})
	}
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) {
	result, ok := obj.(*object.Boolean)
	require.True(t, ok, "expected boolean object")
	require.Equal(t, expected, result.Value, "object's value is wrong")
}

func TestEvalBangOperator(t *testing.T) {
	for _, tt := range []struct {
		source   string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	} {
		t.Run(tt.source, func(t *testing.T) {
			// 1. Act
			got := testEval(t, tt.source)

			// 2. Assert
			testBooleanObject(t, got, tt.expected)
		})
	}
}

func TestEvalIfElse(t *testing.T) {
	for _, tt := range []struct {
		source   string
		expected any
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
	} {
		t.Run(tt.source, func(t *testing.T) {
			// 1. Act
			got := testEval(t, tt.source)

			// 2. Assert
			integer, ok := tt.expected.(int)
			if ok {
				testIntegerObject(t, got, int64(integer))
			} else {
				testNullObject(t, got)
			}
		})
	}
}

func testNullObject(t *testing.T, obj object.Object) {
	require.Equal(t, obj, object.NULL, "object is not NULL. got=%T", obj)
}

func TestReturnStatements(t *testing.T) {
	for _, tt := range []struct {
		source   string
		expected int64
	}{
		{"return 10;", 10},
		{"return 10; 9;", 10},
		{"return 2 * 5; 9;", 10},
		{"9; return 2 * 5; 9;", 10},
		{`
if (10 > 1) {
	if (10 > 1) {
		return 10;
	}
  return 1;
}`,
			10,
		},
	} {
		t.Run(tt.source, func(t *testing.T) {
			// 1. Arrange
			got := testEval(t, tt.source)

			// 2. Assert
			testIntegerObject(t, got, tt.expected)
		})
	}
}

func TestErrorHandling(t *testing.T) {
	for _, tt := range []struct {
		source          string
		expectedMessage string
	}{
		{
			"5 + true;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + true; 5;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-true",
			"unknown operator: -BOOLEAN",
		},
		{
			"true + false;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; true + false; 5",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { true + false; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`
if (10 > 1) {
	if (10 > 1) {
		return true + false;
		}
	return 1;
} `,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
	} {
		t.Run(tt.source, func(t *testing.T) {
			// 1. Act
			got := testEval(t, tt.source)

			// 2. Assert
			result, ok := got.(*object.Error)
			require.True(t, ok, "expected an error object")
			require.Equal(t, tt.expectedMessage, result.Message, "wrong error message")
		})
	}
}
