package expressions

import (
	"interpreter/internal/lexer/tokens"
	"interpreter/internal/object"
)

// Boolean expression that returns a boolean constant: true or false.
// Examples of where a boolean constant can be used below:
//
//	true;
//	false;
//	let x = true;
//	let x = fn() { return true; }
type Boolean struct {
	Token tokens.Token
	Value bool
}

func (b *Boolean) Literal() string {
	return b.Token.Literal
}

func (b *Boolean) String() string {
	return b.Literal()
}

func (b *Boolean) Accept(visitor ExpressionVisitor) object.Object {
	return visitor.VisitBoolean(b)
}

func NewBoolean(token tokens.Token) *Boolean {
	return &Boolean{Token: token, Value: token.Type == tokens.TRUE}
}
