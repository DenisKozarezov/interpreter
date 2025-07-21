package statements

import (
	"bytes"

	"interpreter/internal/ast/expressions"
	"interpreter/internal/lexer/tokens"
	"interpreter/internal/object"
)

type LetStatement struct {
	Token      tokens.Token
	Identifier expressions.Expression
	Value      expressions.Expression
}

func (s *LetStatement) Literal() string {
	return s.Token.Literal
}

func (s *LetStatement) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(s.Literal() + " ")
	buffer.WriteString(s.Identifier.String())
	buffer.WriteString(" = ")

	if s.Value != nil {
		buffer.WriteString(s.Value.String())
	}
	buffer.WriteString(";")
	return buffer.String()
}

func (s *LetStatement) Accept(_ StatementVisitor) object.Object {
	return nil
}
