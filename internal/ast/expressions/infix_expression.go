package expressions

import (
	"bytes"

	"interpreter/internal/ast"
	"interpreter/internal/lexer/tokens"
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
	Operator        string
	LeftExpression  ast.Expression
	RightExpression ast.Expression
}

func (s *InfixExpression) Literal() string {
	return s.Token.Literal
}

func (s *InfixExpression) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("(")
	buffer.WriteString(s.LeftExpression.String())
	buffer.WriteString(" " + s.Operator + " ")
	buffer.WriteString(s.RightExpression.String())
	buffer.WriteString(")")
	return buffer.String()
}

func (s *InfixExpression) expressionNode() {

}
