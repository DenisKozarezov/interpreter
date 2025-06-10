package statements

import "interpreter/internal/lexer/tokens"

// Identifier представляет собой выражение, которое образует строковый идентификатор.
// Это может быть как идентификатор переменной, так и идентификатор функции и т.п. Например:
//
//	let x = 5;
//	    ^
//	let f = myFunc(x, y);
//	    ^    ^^^^  ^  ^
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
