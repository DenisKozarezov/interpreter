package expressions

import (
	"bytes"

	"interpreter/internal/lexer/tokens"
	"interpreter/internal/object"
)

// ConditionExpression represents an expression that contains a Then block and
// an Else block, as well as a Condition to choose between them. If the condition
// is true, then the Then block is triggered, otherwise the code in the Else block
// (if present) will be executed.
//
//	if (x > y) { return x; } else { return y; }
//
// Where:
//  1. 'if' is the keyword.
//  2. '(x > y)' is an infix expression that returns true or false.
//  3. '{ return x; }' defined in the Then execution block if Condition is true.
//  4. 'else' is the keyword of the alternative block.
//  5. '{ return y; }' is an alternative Else block with a false Condition.
type ConditionExpression struct {
	Token     tokens.Token
	Condition Expression
	Then      statement
	Else      statement
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

func (c *ConditionExpression) Accept(visitor ExpressionVisitor) object.Object {
	return visitor.VisitCondition(c)
}
