package expressions

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

func (i *Identifier) Literal() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) expressionNode() {}
