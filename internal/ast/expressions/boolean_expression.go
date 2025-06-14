package expressions

import "interpreter/internal/lexer/tokens"

type Boolean struct {
	Token tokens.Token
	Value bool
}

func (l *Boolean) Literal() string {
	return l.Token.Literal
}

func (l *Boolean) String() string {
	return l.Token.Literal
}

func (l *Boolean) expressionNode() {}
