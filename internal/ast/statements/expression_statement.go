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
	if s.Value == nil {
		return s.Literal()
	} else {
		return s.Value.String()
	}
}

func (s *ExpressionStatement) statementNode() {

}
