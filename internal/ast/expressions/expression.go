package expressions

import (
	"interpreter/internal/ast"
	"interpreter/internal/object"
)

type statement interface {
	ast.Node
}

type Expression interface {
	ast.Node
	Accept(visitor ExpressionVisitor) object.Object
}
