package statements

import (
	"bytes"

	"interpreter/internal/object"
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

func (p *Program) Accept(visitor StatementVisitor) object.Object {
	return visitor.VisitProgram(p)
}
