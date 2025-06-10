package ast

import (
	"bytes"
)

type Program struct {
	Statements []Statement
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
