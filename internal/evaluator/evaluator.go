package evaluator

import (
	"interpreter/internal/ast/expressions"
	"interpreter/internal/ast/statements"
	"interpreter/internal/lexer/tokens"
	"interpreter/internal/object"
)

type IVisitable[TVisitor any, TReturn any] interface {
	Accept(TVisitor) TReturn
}

type (
	VisitableExpression = IVisitable[expressions.ExpressionVisitor, object.Object]
	VisitableStatement  = IVisitable[statements.StatementVisitor, object.Object]
)

func EvaluateExpression(node VisitableExpression) object.Object {
	var v ASTVisitor
	return node.Accept(&v)
}

func EvaluateStatement(node VisitableStatement) object.Object {
	var v ASTVisitor
	return node.Accept(&v)
}

func evalInfixIntegerExpression(left, right object.Object, operator tokens.TokenType) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case tokens.PLUS:
		return &object.Integer{Value: leftVal + rightVal}
	case tokens.MINUS:
		return &object.Integer{Value: leftVal - rightVal}
	case tokens.ASTERISK:
		return &object.Integer{Value: leftVal * rightVal}
	case tokens.SLASH:
		return &object.Integer{Value: leftVal / rightVal}
	case tokens.LT:
		return object.NativeBooleanToObject(leftVal < rightVal)
	case tokens.GT:
		return object.NativeBooleanToObject(leftVal > rightVal)
	case tokens.EQ:
		return object.NativeBooleanToObject(leftVal == rightVal)
	case tokens.NOT_EQ:
		return object.NativeBooleanToObject(leftVal != rightVal)
	default:
		return object.NULL
	}
}

func evalProgram(statements []statements.Statement) object.Object {
	var result object.Object
	for i := range statements {
		result = EvaluateStatement(statements[i])

		if returnValue, ok := result.(*object.Return); ok {
			return returnValue.Value
		}
	}
	return result
}

func evalBlockStatements(statements []statements.Statement) object.Object {
	var result object.Object
	for i := range statements {
		result = EvaluateStatement(statements[i])

		if result != nil && result.Type() == object.RETURN_TYPE {
			return result
		}
	}
	return result
}

func evalBangOperator(right object.Object) object.Object {
	switch right {
	case object.TRUE:
		return object.FALSE
	case object.FALSE:
		return object.TRUE
	case object.NULL:
		return object.TRUE
	default:
		return object.FALSE
	}
}

func evalMinusOperator(right object.Object) object.Object {
	if right.Type() != object.INTEGER_TYPE {
		return object.NULL
	}

	integer := right.(*object.Integer).Value
	return &object.Integer{Value: -integer}
}
