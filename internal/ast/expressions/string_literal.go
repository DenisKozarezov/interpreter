package expressions

import (
	"interpreter/internal/lexer/tokens"
	"interpreter/internal/object"
)

type StringLiteral struct {
	Token tokens.Token
}

func (s *StringLiteral) Literal() string {
	return s.Token.Literal
}

func (s *StringLiteral) String() string {
	return s.Literal()
}

func (s *StringLiteral) Accept(visitor ExpressionVisitor) object.Object {
	return visitor.VisitString(s)
}
