package expressions

import (
	"bytes"
	"strings"

	"interpreter/internal/lexer/tokens"
	"interpreter/internal/object"
)

// FunctionLiteral выражение, которое определяет функцию со списком аргументов Args и блоком
// инструкций Body. Например:
//
//	fn(x, y) { return x + y; }
//
// где:
//  1. 'fn' - ключевое слово;
//  2. 'x, y' - аргументы функции;
//  3. '{ return x + y; }' - блок выполнения, состоящий из одной единственной инструкции statements.ReturnStatement.
type FunctionLiteral struct {
	Token tokens.Token
	Args  []*Identifier
	Body  Statement
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

func (f *FunctionLiteral) Accept(_ ExpressionVisitor) object.Object {
	return nil
}
