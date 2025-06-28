package evaluator

import (
	"interpreter/internal/lexer"
	"interpreter/internal/object"
	"interpreter/internal/parser"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEvalIntegerExpression(t *testing.T) {
	for _, tt := range []struct {
		source   string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
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
	require.Len(t, program.Statements, 1)

	return Evaluate(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) {
	result, ok := obj.(*object.Integer)
	require.True(t, ok, "expected integer object")
	require.Equal(t, expected, result.Value, "object's value is wrong")
}
