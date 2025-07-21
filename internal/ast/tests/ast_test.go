package tests

import (
	"testing"

	"interpreter/internal/ast/expressions"
	"interpreter/internal/ast/statements"
	"interpreter/internal/lexer/tokens"
)

func TestString(t *testing.T) {
	program := &statements.Program{
		Statements: []statements.Statement{
			&statements.LetStatement{
				Token: tokens.NewToken(tokens.LET, "let"),
				Identifier: &expressions.Identifier{
					Token: tokens.NewToken(tokens.IDENTIFIER, "myVar"),
					Value: "myVar",
				},
				Value: &expressions.Identifier{
					Token: tokens.NewToken(tokens.IDENTIFIER, "anotherVar"),
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got = %q", program.String())
	}
}
