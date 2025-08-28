package lexer

import (
	"interpreter/internal/lexer/tokens"
	"strings"
	"testing"
)

func BenchmarkNewTwoCharToken(b *testing.B) {
	// 1. Arrange
	source := `1 << 10;`

	l := NewLexer(strings.NewReader(source))

	// 2. Act
	for b.Loop() {
		l.twoCharToken('<', tokens.L_SHIFT, tokens.LT)
	}
}
