package evaluator

import (
	"interpreter/internal/ast"
	"interpreter/internal/ast/expressions"
	"interpreter/internal/ast/statements"
	"interpreter/internal/lexer/tokens"
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
	case *expressions.InfixExpression:
		return evalInfixExpression(obj)
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
	switch exp.Token.Type {
	case tokens.BANG:
		return evalBangOperator(right)
	case tokens.MINUS:
		return evalMinusOperator(right)
	default:
		return object.NULL
	}
}

func evalInfixExpression(exp *expressions.InfixExpression) object.Object {
	left := Evaluate(exp.LeftExpression)
	right := Evaluate(exp.RightExpression)
	switch {
	case left.Type() == object.INTEGER_TYPE && right.Type() == object.INTEGER_TYPE:
		return evalInfixIntegerExpression(left, right, exp.Token.Type)
	default:
		return object.NULL
	}
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
