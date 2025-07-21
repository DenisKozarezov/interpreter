package expressions

import (
	"bytes"

	"interpreter/internal/lexer/tokens"
	"interpreter/internal/object"
)

// InfixExpression представляет собой выражение, состоящее из двух операндов - левого
// и правого, а также оператора между ними. Классическими примерами такого выражения
// являются следующие виды конструкций:
//
//	a + b;
//	a - b;
//	a / b;
//	a == b;
//	a != b;
//
// и т.п.
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
