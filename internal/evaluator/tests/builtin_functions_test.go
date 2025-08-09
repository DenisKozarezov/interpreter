package tests

import (
	"interpreter/internal/object"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuiltinFunctions(t *testing.T) {
	for _, tt := range []struct {
		source   string
		expected any
	}{
		{`len("")`, 0},
		{`len("four")`, 4},
		{`len("hello world")`, 11},
		{`len(1)`, "argument to 'len' not supported, got INTEGER"},
		{`len("one", "two")`, "wrong number of arguments. got = 2, want = 1"},
	} {
		t.Run(tt.source, func(t *testing.T) {
			// 1. Act
			got := testEval(t, tt.source)

			// 2. Assert
			switch expected := tt.expected.(type) {
			case int:
				testIntegerObject(t, got, int64(expected))
			case string:
				err, ok := got.(*object.Error)
				require.True(t, ok, "expected an error")
				require.Equal(t, tt.expected, err.Message)
			}
		})
	}
}
