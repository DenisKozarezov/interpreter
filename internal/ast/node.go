package ast

import "fmt"

type Node interface {
	fmt.Stringer
	Literal() string
}

type Expression interface {
	Node
}

type Statement interface {
	Node
}
