package expressions

import (
	"bytes"

	"interpreter/internal/lexer/tokens"
	"interpreter/internal/object"
)

// PrefixExpression is an expression consisting of a right operand,
// as well as an operator in front of it. Typical examples:
//
// !a;
// -a;
// --a;
//
// etc.
type PrefixExpression struct {
	Token           tokens.Token
	RightExpression Expression
}

func (s *PrefixExpression) Literal() string {
	return s.Token.Literal
}

func (s *PrefixExpression) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("(")
	buffer.WriteString(s.Token.Literal)
	buffer.WriteString(s.RightExpression.String())
	buffer.WriteString(")")
	return buffer.String()
}

func (s *PrefixExpression) Accept(visitor ExpressionVisitor) object.Object {
	return visitor.VisitPrefix(s)
}
