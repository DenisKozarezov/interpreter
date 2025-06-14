package expressions

import "interpreter/internal/lexer/tokens"

type Boolean struct {
	Token tokens.Token
	Value bool
}

func (b *Boolean) Literal() string {
	return b.Token.Literal
}

func (b *Boolean) String() string {
	return b.Token.Literal
}

func (b *Boolean) expressionNode() {}
