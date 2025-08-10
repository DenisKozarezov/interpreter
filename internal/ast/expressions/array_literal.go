package expressions

import (
	"bytes"
	"interpreter/internal/lexer/tokens"
	"interpreter/internal/object"
	"strings"
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

	items := make([]string, len(a.Items))
	for i := 0; i < len(a.Items); i++ {
		items[i] = a.Items[i].String()
	}
	buffer.WriteString(strings.Join(items, ", "))
	buffer.WriteString("]")
	return buffer.String()
}

func (a *ArrayLiteral) Accept(visitor ExpressionVisitor) object.Object {
	return visitor.VisitArray(a)
}
