package expressions

import (
	"fmt"

	"interpreter/internal/object"
)

type Statement interface {
	fmt.Stringer
}

type Expression interface {
	fmt.Stringer
	Literal() string
	Accept(visitor ExpressionVisitor) object.Object
}
