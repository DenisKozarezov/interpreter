package ast

import (
	"bytes"
	"interpreter/internal/ast/statements"
)

type Program struct {
	Statements []statements.Statement
}

func (p *Program) Literal() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].Literal()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var buffer bytes.Buffer
	for i := range p.Statements {
		buffer.WriteString(p.Statements[i].String())
	}
	return buffer.String()
}
