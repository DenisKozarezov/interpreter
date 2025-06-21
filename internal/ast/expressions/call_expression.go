package expressions

import (
	"bytes"
	"interpreter/internal/ast"
	"interpreter/internal/lexer/tokens"
	"strings"
)

type CallExpression struct {
	Token    tokens.Token
	Function ast.Expression
	Args     []ast.Expression
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

func (c *CallExpression) expressionNode() {}
