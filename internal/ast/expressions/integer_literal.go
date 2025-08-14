package expressions

import (
	"interpreter/internal/lexer/tokens"
	"interpreter/internal/object"
)

// IntegerLiteral is an expression that produces a certain integer constant.
// Consider this as a 'primitive' type for the interpreted scripting language.
// There are so many cases, when objects like integers can be produced in other
// operations as shown below:
//
//	5;
//	let y = 5;
//		    ^
//	if 5 == 5 {
//	   ^   ^
//	if f(x) == f(y) {
//	    ^       ^
//
// As you can see: let instructions, functions, literals - almost every expression
// can return an object, storing an integer constant inside.
type IntegerLiteral struct {
	Token tokens.Token
	Value int64
}

func (l *IntegerLiteral) Literal() string {
	return l.Token.Literal
}

func (l *IntegerLiteral) String() string {
	return l.Literal()
}

func (l *IntegerLiteral) Accept(visitor ExpressionVisitor) object.Object {
	return visitor.VisitInteger(l)
}
