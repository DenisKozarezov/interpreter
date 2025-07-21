package expressions

import "fmt"

type Expression interface {
	fmt.Stringer
	Literal() string
}
