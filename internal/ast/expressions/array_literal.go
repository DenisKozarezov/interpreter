package expressions

import (
	"bytes"
	"interpreter/internal/lexer/tokens"
	"interpreter/internal/object"
)

type ArrayLiteral struct {
	Token tokens.Token
	Items []Expression
}

func (a *ArrayLiteral) Literal() string {
	return a.Token.Literal
}

func (a *ArrayLiteral) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i := 0; i < len(a.Items); i++ {
		buffer.WriteString(a.Items[i].Literal() + ",")
	}
	buffer.WriteString("]")
	return buffer.String()
}

func (a *ArrayLiteral) Accept(_ ExpressionVisitor) object.Object {
	return object.NULL
}
