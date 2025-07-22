package statements

import (
	"interpreter/internal/ast"
	"interpreter/internal/object"
)

type Statement interface {
	ast.Node
	Accept(visitor StatementVisitor) object.Object
}
