package evaluator

import (
	"interpreter/internal/ast"
	"interpreter/internal/ast/expressions"
	"interpreter/internal/ast/statements"
	"interpreter/internal/object"
)

func Evaluate(node ast.Node) object.Object {
	switch obj := node.(type) {
	case *ast.Program:
		return evaluateStatements(obj.Statements)
	case *statements.ExpressionStatement:
		return Evaluate(obj.Value)
	case *expressions.IntegerLiteral:
		return &object.Integer{Value: obj.Value}
	case *expressions.Boolean:
		return object.NativeBooleanToObject(obj.Value)
	case *expressions.PrefixExpression:
		return evalPrefixExpression(obj)
	}
	return nil
}

func evaluateStatements(statements []ast.Statement) object.Object {
	var result object.Object
	for i := range statements {
		result = Evaluate(statements[i])
	}
	return result
}

func evalPrefixExpression(exp *expressions.PrefixExpression) object.Object {
	right := Evaluate(exp.RightExpression)
	switch exp.Operator {
	case "!":
		return evalBangOperator(right)
	case "-":
		return evalMinusOperator(right)
	default:
		return object.NULL
	}
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
		return object.NULL
	}
}

func evalMinusOperator(right object.Object) object.Object {
	if right.Type() != object.INTEGER_TYPE {
		return object.NULL
	}

	integer := right.(*object.Integer).Value
	return &object.Integer{Value: -integer}
}
