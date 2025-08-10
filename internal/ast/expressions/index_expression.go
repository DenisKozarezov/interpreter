package expressions

import (
	"bytes"
	"interpreter/internal/lexer/tokens"
	"interpreter/internal/object"
)

type IndexExpression struct {
	Token          tokens.Token
	LeftExpression Expression
	Index          Expression
}

func (i *IndexExpression) Literal() string {
	return i.Token.Literal
}

func (i *IndexExpression) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("(")
	buffer.WriteString(i.LeftExpression.String())
	buffer.WriteString("[")
	buffer.WriteString(i.Index.String())
	buffer.WriteString("])")
	return buffer.String()
}

func (i *IndexExpression) Accept(visitor ExpressionVisitor) object.Object {
	return visitor.VisitIndex(i)
}
