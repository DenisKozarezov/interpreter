package expressions

import (
	"bytes"
	"interpreter/internal/ast"
	"interpreter/internal/lexer/tokens"
	"strings"
)

// CallExpression представляет собой выражение, которое обозначает вызов функции и содержит
// список аргументов для вызова. Является infix-оператором, поскольку находится между идентификатором
// функции и идентификатором аргумента:
//
//	myFunc(x, y);
//	      ^
//
// где:
//  1. 'myFunc' - идентификатор (см. Identifier) функции Function;
//  2. '(' - оператор вызова функции;
//  3. 'x, y' - идентификаторы аргументов функции.
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
