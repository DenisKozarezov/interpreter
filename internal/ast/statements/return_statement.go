package statements

import (
	"bytes"

	"interpreter/internal/lexer/tokens"
)

type ReturnStatement struct {
	Token tokens.Token
	Value Expression
}

func (s *ReturnStatement) Literal() string {
	return s.Token.Literal
}

func (s *ReturnStatement) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(s.Literal() + " ")

	if s.Value != nil {
		buffer.WriteString(s.Value.String())
	}
	buffer.WriteString(";")
	return buffer.String()
}

func (s *ReturnStatement) statementNode() {

}
