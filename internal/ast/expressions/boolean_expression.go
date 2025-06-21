package expressions

import "interpreter/internal/lexer/tokens"

// Boolean выражение, возвращающее булевую константу: true либо false. Примеры,
// где может применяться булевая константа:
//
//	true;
//	false;
//	let x = true;
//	let x = fn() { return true; }
type Boolean struct {
	Token tokens.Token
	Value bool
}

func (b *Boolean) Literal() string {
	return b.Token.Literal
}

func (b *Boolean) String() string {
	return b.Literal()
}

func (b *Boolean) expressionNode() {}

func NewBoolean(token tokens.Token) *Boolean {
	return &Boolean{Token: token, Value: token.Type == tokens.TRUE}
}
