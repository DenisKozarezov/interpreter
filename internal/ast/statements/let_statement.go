package statements

import (
	"interpreter/internal/ast"
	"interpreter/internal/lexer/tokens"
)

type LetStatement struct {
	Token      tokens.Token
	Identifier *Identifier
	Value      ast.Expression
}

func (l *LetStatement) Literal() string {
	return l.Token.Literal
}

func (l *LetStatement) String() string {
	return l.Token.Literal
}

func (l *LetStatement) statementNode() {

}
