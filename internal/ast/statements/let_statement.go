package statements

import (
	"bytes"

	"interpreter/internal/ast"
	"interpreter/internal/lexer/tokens"
)

type LetStatement struct {
	Token      tokens.Token
	Identifier ast.Expression
	Value      ast.Expression
}

func (s *LetStatement) Literal() string {
	return s.Token.Literal
}

func (s *LetStatement) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(s.Literal() + " ")
	buffer.WriteString(s.Identifier.String())
	buffer.WriteString(" = ")

	if s.Value != nil {
		buffer.WriteString(s.Value.String())
	}
	buffer.WriteString(";")
	return buffer.String()
}

func (s *LetStatement) statementNode() {

}
