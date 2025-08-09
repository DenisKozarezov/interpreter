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

func EvaluateExpression(node VisitableExpression, visitor expressions.ExpressionVisitor) object.Object {
	if node == nil {
		return newRuntimeError("AST expression node is nil")
	}

	return node.Accept(visitor)
}

func EvaluateStatement(node VisitableStatement, visitor statements.StatementVisitor) object.Object {
	if node == nil {
		return newRuntimeError("AST statement node is nil")
	}

	return node.Accept(visitor)
}

func evalInfixIntegerExpression(left, right object.Object, operator tokens.Token) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator.Type {
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
		return newRuntimeError("unknown operator: %s %s %s", left.Type(), operator.Literal, right.Type())
	}
}

func evalInfixStringExpression(left, right object.Object, operator tokens.Token) object.Object {
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value

	switch operator.Type {
	case tokens.PLUS:
		return &object.String{Value: leftVal + rightVal}
	default:
		return newRuntimeError("unknown operator: %s %s %s", left.Type(), operator.Literal, right.Type())
	}
}

func evalProgram(statements []statements.Statement, visitor statements.StatementVisitor) object.Object {
	var result object.Object

	for i := range statements {
		result = EvaluateStatement(statements[i], visitor)

		switch obj := result.(type) {
		case *object.Return:
			return obj.Value
		case *object.Error:
			return obj
		}
	}
	return result
}

func evalBlockStatements(statements []statements.Statement, visitor statements.StatementVisitor) object.Object {
	var result object.Object
	for i := range statements {
		result = EvaluateStatement(statements[i], visitor)

		if result != nil && result.Type() == object.RETURN_TYPE || isRuntimeError(result) {
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
		return newRuntimeError("unknown operator: -%s", right.Type())
	}

	integer := right.(*object.Integer).Value
	return &object.Integer{Value: -integer}
}

func evalExpressions(expressions []expressions.Expression, visitor expressions.ExpressionVisitor) []object.Object {
	result := make([]object.Object, len(expressions))
	for i := 0; i < len(expressions); i++ {
		val := EvaluateExpression(expressions[i], visitor)
		if isRuntimeError(val) {
			return []object.Object{val}
		}
		result[i] = val
	}
	return result
}

func extendFunctionEnvironment(fn *object.Function, args []object.Object) *object.Environment {
	enclosedEnv := object.NewEnclosedEnvironment(fn.Environment)

	for i := 0; i < len(fn.Args); i++ {
		enclosedEnv.Set(fn.Args[i].String(), args[i])
	}

	return enclosedEnv
}

func applyFunction(fnBody statements.Statement, visitor statements.StatementVisitor) object.Object {
	return unwrapReturnValue(EvaluateStatement(fnBody, visitor))
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.Return); ok {
		return returnValue
	}
	return obj
}
