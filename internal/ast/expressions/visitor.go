package expressions

import "interpreter/internal/object"

type ExpressionVisitor interface {
	VisitInteger(integer *IntegerLiteral) object.Object
	VisitBoolean(boolean *Boolean) object.Object
	VisitPrefix(prefix *PrefixExpression) object.Object
	VisitInfix(infix *InfixExpression) object.Object
	VisitCondition(condition *ConditionExpression) object.Object
	VisitIdentifier(identifier *Identifier) object.Object
	VisitFunction(function *FunctionLiteral) object.Object
	VisitCallExpression(call *CallExpression) object.Object
}
