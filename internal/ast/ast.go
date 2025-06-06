package ast

import "fmt"

type Node interface {
	fmt.Stringer
	Literal() string
}

type Statement interface {
	Node
}

type Expression interface {
	Node
}

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
