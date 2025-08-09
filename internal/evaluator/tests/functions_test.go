package tests

import (
	"interpreter/internal/object"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFunctionObject(t *testing.T) {
	// 1. Arrange
	source := "fn(x) { x + 2; };"

	// 2. Act
	got := testEval(t, source)

	// 3. Assert
	fn, ok := got.(*object.Function)
	require.True(t, ok, "expected function, but got=%T", got)
	require.Len(t, fn.Args, 1, "expected one parameter in function")
	require.Equal(t, "x", fn.Args[0].String(), "parameter is not x")
	require.Equal(t, "(x + 2)", fn.Body.String())
}

func TestFunctionEvaluation(t *testing.T) {
	for _, tt := range []struct {
		source   string
		expected int64
	}{
		{"let identity = fn(x) { x; }; identity(5);", 5},
		{"let identity = fn(x) { return x; }; identity(5);", 5},
		{"let double = fn(x) { x * 2; }; double(5);", 10},
		{"let add = fn(x, y) { x + y; }; add(5, 5);", 10},
		{"let add = fn(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"fn(x) { x; }(5)", 5},
	} {
		t.Run(tt.source, func(t *testing.T) {
			// 1. Arrange
			got := testEval(t, tt.source)

			// 2. Act
			testIntegerObject(t, got, tt.expected)
		})
	}
}

func TestClosures(t *testing.T) {
	// 1. Arrange
	source := `
let newAdder = fn(x) {
	fn(y) { x + y };
};

let addTwo = newAdder(2);
addTwo(2);
`

	// 2. Act
	got := testEval(t, source)

	// 3. Assert
	testIntegerObject(t, got, 4)
}
