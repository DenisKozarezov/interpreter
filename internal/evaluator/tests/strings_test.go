package tests

import (
	"interpreter/internal/object"
	"testing"

	"github.com/stretchr/testify/require"
)

func testStringLiteral(t *testing.T, obj object.Object, expectedString string) {
	str, ok := obj.(*object.String)
	require.True(t, ok, "expected a string")
	require.Equalf(t, expectedString, str.Value, "expected literal '%s'", expectedString)
}

func TestEvaluateStringLiterals(t *testing.T) {
	// 1. Arrange
	source := `"hello world"`

	// 2. Act
	got := testEval(t, source)

	// 3. Assert
	testStringLiteral(t, got, "hello world")
}

func TestStringConcatenation(t *testing.T) {
	for _, tt := range []struct {
		source   string
		expected string
	}{
		{
			source:   `""`,
			expected: "",
		},
		{
			source:   `"" + ""`,
			expected: "",
		},
		{
			source:   `"" + " " + ""`,
			expected: " ",
		},
		{
			source:   `"hello" + " " + "world"`,
			expected: "hello world",
		},
		{
			source:   `"1" + "2" + "3" + "4" + "5"`,
			expected: "12345",
		},
		{
			source:   `"1"+"2"+"3"+"4"+"5"`,
			expected: "12345",
		},
	} {
		t.Run(tt.source, func(t *testing.T) {
			// 1. Act
			got := testEval(t, tt.source)

			// 2. Assert
			testStringLiteral(t, got, tt.expected)
		})
	}
}
