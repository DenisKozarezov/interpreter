package expressions

import (
	"bytes"
	"strings"

	"interpreter/internal/lexer/tokens"
	"interpreter/internal/object"
)

// CallExpression is an expression that denotes a function call and contains
// a list of arguments for the call. It is an infix operator because it is
// located between the function identifier and the argument identifier:
//
//	myFunc(x, y);
//	      ^
//
// Where:
//  1. 'myFunc' is the identifier (see Identifier) of the function.
//  2. '(' is the function call operator.
//  3. 'x, y' are the argument identifiers of the function.
type CallExpression struct {
	Token    tokens.Token
	Function Expression
	Args     []Expression
}

func (c *CallExpression) Literal() string {
	return c.Token.Literal
}

func (c *CallExpression) String() string {
	var buffer bytes.Buffer

	args := make([]string, len(c.Args))
	for i := range c.Args {
		args[i] = c.Args[i].String()
	}

	buffer.WriteString(c.Function.String())
	buffer.WriteString("(")
	buffer.WriteString(strings.Join(args, ", "))
	buffer.WriteString(")")
	return buffer.String()
}

func (c *CallExpression) Accept(visitor ExpressionVisitor) object.Object {
	return visitor.VisitCallExpression(c)
}
