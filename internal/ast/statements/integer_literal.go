package statements

import (
	"interpreter/internal/lexer/tokens"
)

type IntegerLiteral struct {
	Token tokens.Token
	Value int64
}

func (l *IntegerLiteral) Literal() string {
	return l.Token.Literal
}

func (l *IntegerLiteral) String() string {
	return l.Token.Literal
}

func (l *IntegerLiteral) expressionNode() {}
