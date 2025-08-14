package expressions

import (
	"interpreter/internal/lexer/tokens"
	"interpreter/internal/object"
)

// Identifier is an expression that forms a string literal associated to some data
// in heap or stack. It can be a variable identifier, a function identifier, whatever...
// You should always remember: identifiers do not hold a value. Only names.
//
// For example:
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
