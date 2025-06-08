package ast

import (
	"interpreter/internal/ast/statements"
	"interpreter/internal/lexer/tokens"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []statements.Statement{
			&statements.LetStatement{
				Token: tokens.NewToken(tokens.LET, "let"),
				Identifier: &statements.Identifier{
					Token: tokens.NewToken(tokens.IDENTIFIER, "myVar"),
					Value: "myVar",
				},
				Value: &statements.Identifier{
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
