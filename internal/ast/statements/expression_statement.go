package statements

import (
	"interpreter/internal/ast"
	"interpreter/internal/lexer/tokens"
)

type ExpressionStatement struct {
	Token tokens.Token
	Value ast.Expression
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

func (s *ExpressionStatement) statementNode() {

}

func NewStatement(token tokens.Token, exp ast.Expression) *ExpressionStatement {
	return &ExpressionStatement{Token: token, Value: exp}
}
