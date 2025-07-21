package statements

import (
	"bytes"

	"interpreter/internal/lexer/tokens"
)

// BlockStatement представляет собой блок различных инструкций. Например:
//
//  1. Пустое множество: {  }
//
//  2. Одна инструкция: { return 1; }
//
//  3. Несколько инструкций:
//
//     { let x = 1; return x; }
type BlockStatement struct {
	Token      tokens.Token
	Statements []Statement
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
