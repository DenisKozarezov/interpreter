package expressions

import (
	"bytes"

	"interpreter/internal/lexer/tokens"
	"interpreter/internal/object"
)

// InfixExpression is an expression consisting of two operands, the left one and the right one,
// and an operator between them. Typical examples below:
//
//	a + b;
//	a - b;
//	a / b;
//	a == b;
//	a != b;
//
// etc.
type InfixExpression struct {
	Token           tokens.Token
	LeftExpression  Expression
	RightExpression Expression
}

func (s *InfixExpression) Literal() string {
	return s.Token.Literal
}

func (s *InfixExpression) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("(")
	buffer.WriteString(s.LeftExpression.String())
	buffer.WriteString(" " + s.Token.Literal + " ")
	buffer.WriteString(s.RightExpression.String())
	buffer.WriteString(")")
	return buffer.String()
}

func (s *InfixExpression) Accept(visitor ExpressionVisitor) object.Object {
	return visitor.VisitInfix(s)
}
