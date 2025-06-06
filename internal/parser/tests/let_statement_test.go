package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"interpreter/internal/ast"
	"interpreter/internal/ast/statements"
	lex "interpreter/internal/lexer"
	"interpreter/internal/lexer/tokens"
	"interpreter/internal/parser"
)

func TestLetStatement(t *testing.T) {
	source := `
let x = 5;
let y = 10;
let foobar = 838383;
let = 15;
`
	// 1. Arrange
	l := lex.NewLexer(source)
	p := parser.NewParser(l)

	// 2. Act
	program := p.Parse()

	// 3. Assert
	require.Len(t, program.Statements, 3)
	require.Len(t, p.Errors(), 1)

	for i, tt := range []struct {
		name               string
		expectedIdentifier string
	}{
		{"x", "x"},
		{"y", "y"},
		{"foobar", "foobar"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ok, err := checkLetStatement(program.Statements[i], tt.expectedIdentifier)
			if !ok {
				t.Fatal(err)
				return
			}
		})
	}
}

func checkLetStatement(s ast.Statement, expectedIdentifier string) (bool, error) {
	tokenType := tokens.LookupIdentifierType(s.Literal())
	if tokenType != tokens.LET {
		return false, fmt.Errorf("expected let literal, got %s [%d]", s.Literal(), tokenType)
	}

	statement, ok := s.(*statements.LetStatement)
	if !ok {
		return false, fmt.Errorf("expected let statement, got %s", statement.Literal())
	}

	if statement.Identifier.Literal() != expectedIdentifier {
		return false, fmt.Errorf("expected identifier literal %s, got %s", expectedIdentifier, s.Literal())
	}

	if statement.Identifier.Value != expectedIdentifier {
		return false, fmt.Errorf("expected identifier value %s, got %s", expectedIdentifier, s.Literal())
	}

	return true, nil
}
