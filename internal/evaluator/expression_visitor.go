package evaluator

import (
	"interpreter/internal/ast/expressions"
	"interpreter/internal/ast/statements"
	"interpreter/internal/lexer/tokens"
	"interpreter/internal/object"
)

type ASTVisitor struct {
	env *object.Environment
}

func NewASTVisitor() *ASTVisitor {
	return &ASTVisitor{env: object.NewEnvironment()}
}

func (v *ASTVisitor) VisitInteger(integer *expressions.IntegerLiteral) object.Object {
	return &object.Integer{Value: integer.Value}
}

func (v *ASTVisitor) VisitBoolean(boolean *expressions.Boolean) object.Object {
	return object.NativeBooleanToObject(boolean.Value)
}

func (v *ASTVisitor) VisitPrefix(prefix *expressions.PrefixExpression) object.Object {
	right := EvaluateExpression(prefix.RightExpression, v)
	if isRuntimeError(right) {
		return right
	}

	switch prefix.Token.Type {
	case tokens.BANG:
		return evalBangOperator(right)
	case tokens.MINUS:
		return evalMinusOperator(right)
	default:
		return newRuntimeError("unknown operator: %s%s", prefix.Token.Literal, right.Type())
	}
}

func (v *ASTVisitor) VisitInfix(infix *expressions.InfixExpression) object.Object {
	left := EvaluateExpression(infix.LeftExpression, v)
	if isRuntimeError(left) {
		return left
	}

	right := EvaluateExpression(infix.RightExpression, v)
	if isRuntimeError(right) {
		return right
	}

	switch {
	case left.Type() == object.INTEGER_TYPE && right.Type() == object.INTEGER_TYPE:
		return evalInfixIntegerExpression(left, right, infix.Token)
	case left.Type() != right.Type():
		return newRuntimeError("type mismatch: %s %s %s", left.Type(), infix.Token.Literal, right.Type())
	default:
		return newRuntimeError("unknown operator: %s %s %s", left.Type(), infix.Token.Literal, right.Type())
	}
}

func (v *ASTVisitor) VisitCondition(condition *expressions.ConditionExpression) object.Object {
	cond := EvaluateExpression(condition.Condition, v)
	if isRuntimeError(cond) {
		return cond
	}

	if object.ObjectToNativeBoolean(cond) {
		return EvaluateStatement(condition.Then.(*statements.BlockStatement), v)
	} else if condition.Else != nil {
		return EvaluateStatement(condition.Else.(*statements.BlockStatement), v)
	} else {
		return object.NULL
	}
}

func (v *ASTVisitor) VisitIdentifier(identifier *expressions.Identifier) object.Object {
	name := identifier.Literal()
	val, ok := v.env.Get(name)
	if !ok {
		return newRuntimeError("identifier not found: %s", name)
	}

	return val
}
