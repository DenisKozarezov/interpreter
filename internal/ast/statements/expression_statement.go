package statements

import (
	"interpreter/internal/ast"
	"interpreter/internal/lexer/tokens"
)

type ExpressionStatement struct {
	Token tokens.Token
	Value ast.Expression
}

func (l *ExpressionStatement) Literal() string {
	return l.Token.Literal
}

func (l *ExpressionStatement) String() string {
	return l.Token.Literal
}

func (l *ExpressionStatement) statementNode() {

}
