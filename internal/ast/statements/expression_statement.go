package statements

import (
	"interpreter/internal/lexer/tokens"
)

type ExpressionStatement struct {
	Token tokens.Token
	Value Expression
}

func (s *ExpressionStatement) Literal() string {
	return s.Token.Literal
}

func (s *ExpressionStatement) String() string {
	return s.Token.Literal
}

func (s *ExpressionStatement) statementNode() {

}
