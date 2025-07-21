package evaluator

import (
	"interpreter/internal/ast/expressions"
	"interpreter/internal/ast/statements"
	"interpreter/internal/lexer/tokens"
	"interpreter/internal/object"
)

type ASTVisitor struct{}

func (v *ASTVisitor) VisitProgram(program *statements.Program) object.Object {
	var result object.Object
	for i := range program.Statements {
		result = EvaluateStatement(program.Statements[i])
	}
	return result
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
	switch prefix.Token.Type {
	case tokens.BANG:
		return evalBangOperator(right)
	case tokens.MINUS:
		return evalMinusOperator(right)
	default:
		return object.NULL
	}
}

func (v *ASTVisitor) VisitInfix(infix *expressions.InfixExpression) object.Object {
	left := EvaluateExpression(infix.LeftExpression)
	right := EvaluateExpression(infix.RightExpression)
	switch {
	case left.Type() == object.INTEGER_TYPE && right.Type() == object.INTEGER_TYPE:
		return evalInfixIntegerExpression(left, right, infix.Token.Type)
	default:
		return object.NULL
	}
}
