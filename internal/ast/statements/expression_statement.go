package statements

import (
	"interpreter/internal/ast/expressions"
	"interpreter/internal/lexer/tokens"
	"interpreter/internal/object"
)

type ExpressionStatement struct {
	Token tokens.Token
	Value expressions.Expression
}

func (s *ExpressionStatement) Literal() string {
	return s.Token.Literal
}

func (s *ExpressionStatement) String() string {
	if s.Value == nil {
		return s.Literal()
	} else {
		return s.Value.String()
	}
}

func (s *ExpressionStatement) Accept(visitor StatementVisitor) object.Object {
	return visitor.VisitExpression(s)
}

func NewStatement(token tokens.Token, exp expressions.Expression) *ExpressionStatement {
	return &ExpressionStatement{Token: token, Value: exp}
}
