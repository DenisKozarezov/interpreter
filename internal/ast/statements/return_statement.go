package statements

import (
	"interpreter/internal/ast"
	"interpreter/internal/lexer/tokens"
)

type ReturnStatement struct {
	Token tokens.Token
	Value ast.Expression
}

func (l *ReturnStatement) Literal() string {
	return l.Token.Literal
}

func (l *ReturnStatement) String() string {
	return l.Token.Literal
}

func (l *ReturnStatement) statementNode() {

}
