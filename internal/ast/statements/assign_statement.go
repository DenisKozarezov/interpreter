package statements

import (
	"bytes"
	"interpreter/internal/ast/expressions"
	"interpreter/internal/lexer/tokens"
	"interpreter/internal/object"
)

type AssignStatement struct {
	Token      tokens.Token
	Identifier tokens.Token
	Expression expressions.Expression
}

func (s *AssignStatement) Literal() string {
	return s.Identifier.Literal
}

func (s *AssignStatement) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(s.Identifier.Literal)
	buffer.WriteString(" = ")

	if s.Expression != nil {
		buffer.WriteString(s.Expression.String())
	}
	buffer.WriteString(";")

	return buffer.String()
}

func (s *AssignStatement) Accept(visitor StatementVisitor) object.Object {
	return visitor.VisitAssign(s)
}
