package expressions

import (
	"bytes"
	"interpreter/internal/ast"
	"interpreter/internal/lexer/tokens"
)

// ConditionExpression обозначает выражение, содержажее блок выполнения Then и
// блок альтернативы Else, а также условие Condition для выбора между ними. Если условие
// истинно, то срабатывает блок Then, в противном случае выполняться будет фрагмент кода
// в блоке Else (если он присутствует).
//
//	if (x > y) { return x; } else { return y; }
//
// где:
//  1. 'if' - ключевое слово;
//  2. '(x > y)' - infix-выражение, возвращающее истину либо ложь;
//  3. '{ return x; }' - блок выполнения Then при истинности условия Condition;
//  4. 'else' - ключевое слово блока альтернативы;
//  5. '{ return y; }' - блок альтернативы Else при ложном условии Condition.
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
