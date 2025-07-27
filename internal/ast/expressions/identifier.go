package expressions

import (
	"interpreter/internal/lexer/tokens"
	"interpreter/internal/object"
)

// Identifier представляет собой выражение, которое образует строковый идентификатор.
// Это может быть как идентификатор переменной, так и идентификатор функции и т.п. Например:
//
//	let x = 5;
//	    ^
//	let f = myFunc(x, y);
//	    ^    ^^^^  ^  ^
type Identifier struct {
	Token tokens.Token
	Value string
}

func (i *Identifier) Literal() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) Accept(visitor ExpressionVisitor) object.Object {
	return visitor.VisitIdentifier(i)
}

func NewIdentifier(token tokens.Token) *Identifier {
	return &Identifier{Token: token, Value: token.Literal}
}
