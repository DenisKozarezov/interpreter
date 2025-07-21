package statements

import (
	"fmt"

	"interpreter/internal/object"
)

type Statement interface {
	fmt.Stringer
	Literal() string
	Accept(visitor StatementVisitor) object.Object
}
