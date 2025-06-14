package expressions

import (
	"bytes"
	"interpreter/internal/ast"
	"interpreter/internal/lexer/tokens"
)

type BlockStatement struct {
	Token      tokens.Token
	Statements []ast.Statement
}

func (s *BlockStatement) Literal() string {
	return s.Token.Literal
}

func (s *BlockStatement) String() string {
	var buffer bytes.Buffer
	for i := range s.Statements {
		buffer.WriteString(s.Statements[i].String())
	}
	return buffer.String()
}

func (s *BlockStatement) statementNode() {

}
