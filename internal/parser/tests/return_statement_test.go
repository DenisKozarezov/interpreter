package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"interpreter/internal/ast/statements"
	lex "interpreter/internal/lexer"
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
	l := lex.NewLexer(source)
	p := parser.NewParser(l)

	// 2. Act
	program := p.Parse()

	// 3. Assert
	require.Len(t, program.Statements, 5)
	require.Zero(t, p.Errors())

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
			ok, err := checkReturnStatement(program.Statements[i])
			if !ok {
				t.Fatal(err)
				return
			}
		})
	}
}

func checkReturnStatement(s statements.Statement) (bool, error) {
	tokenType := tokens.LookupIdentifierType(s.Literal())
	if tokenType != tokens.RETURN {
		return false, fmt.Errorf("expected return literal, got %s [%d]", s.Literal(), tokenType)
	}

	statement, ok := s.(*statements.ReturnStatement)
	if !ok {
		return false, fmt.Errorf("expected return statement, got %s", statement.Literal())
	}

	return true, nil
}
