package evaluator

import (
	"interpreter/internal/ast/expressions"
	"interpreter/internal/ast/statements"
	"interpreter/internal/lexer/tokens"
	"interpreter/internal/object"
)

type ASTVisitor struct{}

func (v *ASTVisitor) VisitProgram(program *statements.Program) object.Object {
	return evalProgram(program.Statements)
}

func (v *ASTVisitor) VisitExpression(expression *statements.ExpressionStatement) object.Object {
	return EvaluateExpression(expression.Value)
}

func (v *ASTVisitor) VisitInteger(integer *expressions.IntegerLiteral) object.Object {
	return &object.Integer{Value: integer.Value}
}

func (v *ASTVisitor) VisitBoolean(boolean *expressions.Boolean) object.Object {
	return object.NativeBooleanToObject(boolean.Value)
}

func (v *ASTVisitor) VisitPrefix(prefix *expressions.PrefixExpression) object.Object {
	right := EvaluateExpression(prefix.RightExpression)
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
	left := EvaluateExpression(infix.LeftExpression)
	if isRuntimeError(left) {
		return left
	}

	right := EvaluateExpression(infix.RightExpression)
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
	cond := EvaluateExpression(condition.Condition)
	if isRuntimeError(cond) {
		return cond
	}

	if object.ObjectToNativeBoolean(cond) {
		return EvaluateStatement(condition.Then.(*statements.BlockStatement))
	} else if condition.Else != nil {
		return EvaluateStatement(condition.Else.(*statements.BlockStatement))
	} else {
		return object.NULL
	}
}

func (v *ASTVisitor) VisitBlockStatement(block *statements.BlockStatement) object.Object {
	return evalBlockStatements(block.Statements)
}

func (v *ASTVisitor) VisitReturn(r *statements.ReturnStatement) object.Object {
	value := EvaluateExpression(r.Value)
	if isRuntimeError(value) {
		return value
	}
	return &object.Return{Value: value}
}

func (v *ASTVisitor) VisitLetStatement(let *statements.LetStatement) object.Object {
	value := EvaluateExpression(let.Value)
	if isRuntimeError(value) {
		return value
	}

	// TODO: продумать контекст (environment) в текущем блоке
	return object.NULL
}
