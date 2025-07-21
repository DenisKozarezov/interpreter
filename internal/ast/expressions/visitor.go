package expressions

import "interpreter/internal/object"

type ExpressionVisitor interface {
	VisitInteger(integer *IntegerLiteral) object.Object
	VisitBoolean(boolean *Boolean) object.Object
	VisitPrefix(prefix *PrefixExpression) object.Object
	VisitInfix(infix *InfixExpression) object.Object
}
