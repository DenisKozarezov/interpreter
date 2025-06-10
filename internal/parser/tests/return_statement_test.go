package tests

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"interpreter/internal/ast/statements"
	"interpreter/internal/lexer"
	"interpreter/internal/lexer/tokens"
	"interpreter/internal/parser"
)

func TestReturnStatement(t *testing.T) {
	source := `
return 5;
return 10;
return x;
return add(5, 10);
return add(x, x);
`

	// 1. Arrange
	l := lexer.NewLexer(strings.NewReader(source))
	p := parser.NewParser(l)

	// 2. Act
	program := p.Parse()

	// 3. Assert
	require.Zero(t, p.Errors())
	require.Len(t, program.Statements, 5)

	for i, tt := range []struct {
		name string
	}{
		{"return 5"},
		{"return 10"},
		{"return x"},
		{"return add(5, 10)"},
		{"return add(x, x)"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			tokenType := tokens.LookupIdentifierType(program.Statements[i].Literal())
			require.Equal(t, tokens.RETURN, tokenType, "expected return literal")

			_, ok := program.Statements[i].(*statements.ReturnStatement)
			require.True(t, ok, "expected return statement")
		})
	}
}
