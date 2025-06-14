package expressions

import (
	"bytes"
	"interpreter/internal/ast"
	"interpreter/internal/lexer/tokens"
)

type ConditionExpression struct {
	Token     tokens.Token
	Condition ast.Expression
	Then      *BlockStatement
	Else      *BlockStatement
}

func (c *ConditionExpression) Literal() string {
	return c.Token.Literal
}

func (c *ConditionExpression) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("if ")
	buffer.WriteString(c.Condition.String())
	buffer.WriteString(" { ")
	buffer.WriteString(c.Then.String())
	buffer.WriteString(" } ")
	if c.Else != nil {
		buffer.WriteString("else { ")
		buffer.WriteString(c.Else.String())
		buffer.WriteString(" } ")
	}
	return buffer.String()
}

func (c *ConditionExpression) expressionNode() {}
