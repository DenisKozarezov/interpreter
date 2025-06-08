package statements

import "interpreter/internal/lexer/tokens"

type Identifier struct {
	Token tokens.Token
	Value string
}

func (l *Identifier) Literal() string {
	return l.Token.Literal
}

func (l *Identifier) String() string {
	return l.Value
}

func (l *Identifier) expressionNode() {}
