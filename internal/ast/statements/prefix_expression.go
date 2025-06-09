package statements

import (
	"bytes"

	"interpreter/internal/lexer/tokens"
)

type PrefixExpression struct {
	Token           tokens.Token
	Operator        string
	RightExpression Expression
}

func (s *PrefixExpression) Literal() string {
	return s.Token.Literal
}

func (s *PrefixExpression) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("(")
	buffer.WriteString(s.Operator)
	buffer.WriteString(s.RightExpression.String())
	buffer.WriteString(")")
	return buffer.String()
}

func (s *PrefixExpression) expressionNode() {

}
