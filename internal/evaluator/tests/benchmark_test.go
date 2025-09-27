package tests

import (
	"testing"
)

func BenchmarkIntegers(b *testing.B) {
	// 1. Arrange
	for _, tt := range []struct {
		source string
	}{
		{"5"},
		{"5 + 5"},
		{"5 + 5 - 5 * 5"},
		{"5 + 5 - 5 * 5 - 10"},
		{"5 + 10"},
	} {
		b.Run(tt.source, func(b *testing.B) {
			// 1. Act
			for b.Loop() {
				testEval(b, tt.source)
			}
		})
	}
}

func BenchmarkStrings(b *testing.B) {
	// 1. Arrange
	for _, tt := range []struct {
		source string
	}{
		{`""`},
		{`" "`},
		{`"123"`},
		{`"123" + "123"`},
		{`"123" + "456"`},
	} {
		b.Run(tt.source, func(b *testing.B) {
			// 1. Act
			for b.Loop() {
				testEval(b, tt.source)
			}
		})
	}
}
