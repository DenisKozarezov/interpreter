package statements

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
