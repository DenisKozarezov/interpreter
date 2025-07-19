package expressions

import (
	"bytes"

	"interpreter/internal/ast"
	"interpreter/internal/lexer/tokens"
)

// PrefixExpression представляет собой выражение, состоящее из правого операнда,
// а также оператора перед ним. Классическими примерами такого выражения являются
// следующие виды конструкций:
//
//	!a;
//	-a;
//	--a;
//
// и т.п.
type PrefixExpression struct {
	Token           tokens.Token
	RightExpression ast.Expression
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

func (s *PrefixExpression) expressionNode() {

}
