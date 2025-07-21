package statements

import (
	"interpreter/internal/object"
)

type StatementVisitor interface {
	VisitProgram(program *Program) object.Object
	VisitExpression(boolean *ExpressionStatement) object.Object
}
