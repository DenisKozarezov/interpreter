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

func TestLetStatement(t *testing.T) {
	source := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	// 1. Arrange
	l := lexer.NewLexer(strings.NewReader(source))
	p := parser.NewParser(l)

	// 2. Act
	program := p.Parse()

	// 3. Assert
	require.Len(t, p.Errors(), 0)
	require.Len(t, program.Statements, 3)

	for i, tt := range []struct {
		name               string
		expectedIdentifier string
	}{
		{"x", "x"},
		{"y", "y"},
		{"foobar", "foobar"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			tokenType := tokens.LookupIdentifierType(program.Statements[i].Literal())
			require.Equal(t, tokens.LET, tokenType, "expected let literal")

			statement, ok := program.Statements[i].(*statements.LetStatement)
			require.True(t, ok, "expected let statement")
			require.Equal(t, tt.expectedIdentifier, statement.Identifier.Literal(), "expected identifier literal")
			require.Equal(t, tt.expectedIdentifier, statement.Identifier.Value, "expected identifier value")
		})
	}
}
