package statements

import (
	"bytes"

	"interpreter/internal/lexer/tokens"
)

type InfixExpression struct {
	Token           tokens.Token
	Operator        string
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
	buffer.WriteString(" " + s.Operator + " ")
	buffer.WriteString(s.RightExpression.String())
	buffer.WriteString(")")
	return buffer.String()
}

func (s *InfixExpression) expressionNode() {

}
