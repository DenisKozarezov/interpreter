package statements

import (
	"interpreter/internal/lexer/tokens"
)

// IntegerLiteral представляет собой выражение, которое производит на свет некую
// целочисленную константу. Например:
//
//	5;
//	^
//
// Важно понимать, что само по себе число 5 является лишь РЕЗУЛЬТАТОМ выражения, а
// не самим выражением. Это необходимо, чтобы были валидны следующие конструкции:
//
//	let y = 5;
//	if 5 == 5 {
//	   ^    ^
type IntegerLiteral struct {
	Token tokens.Token
	Value int64
}

func (l *IntegerLiteral) Literal() string {
	return l.Token.Literal
}

func (l *IntegerLiteral) String() string {
	return l.Literal()
}

func (l *IntegerLiteral) expressionNode() {}
