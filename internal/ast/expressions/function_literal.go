package expressions

import (
	"bytes"
	"strings"

	"interpreter/internal/lexer/tokens"
	"interpreter/internal/object"
)

// FunctionLiteral is an expression that defines a function with an argument list Args
// and a block of instructions Body. For example:
//
//	fn(x, y) { return x + y; }
//
// Where:
// 1. 'fn' is the keyword.
// 2. 'x, y' are the function arguments.
// 3. '{ return x + y; }' is a block (body) of instructions.
type FunctionLiteral struct {
	Token tokens.Token
	Args  []*Identifier
	Body  statement
}

func (f *FunctionLiteral) Literal() string {
	return f.Token.Literal
}

func (f *FunctionLiteral) String() string {
	var buffer bytes.Buffer

	args := make([]string, len(f.Args))
	for i := range f.Args {
		args[i] = f.Args[i].String()
	}

	buffer.WriteString(f.Literal())
	buffer.WriteString("(")
	buffer.WriteString(strings.Join(args, ", "))
	buffer.WriteString(") {")
	buffer.WriteString(f.Body.String())
	buffer.WriteString("}")
	return buffer.String()
}

func (f *FunctionLiteral) Accept(visitor ExpressionVisitor) object.Object {
	return visitor.VisitFunction(f)
}
