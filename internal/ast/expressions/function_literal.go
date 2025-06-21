package expressions

import (
	"bytes"
	"interpreter/internal/lexer/tokens"
	"strings"
)

type FunctionLiteral struct {
	Token tokens.Token
	Args  []*Identifier
	Body  *BlockStatement
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
	buffer.WriteString(strings.Join(args, ","))
	buffer.WriteString(") {")
	buffer.WriteString(f.Body.String())
	buffer.WriteString("}")
	return buffer.String()
}

func (f *FunctionLiteral) expressionNode() {}
