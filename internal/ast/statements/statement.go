package statements

import "fmt"

type Expression interface {
	fmt.Stringer
	Literal() string
}

type Statement interface {
	fmt.Stringer
	Literal() string
}
