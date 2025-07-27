package evaluator

import (
	"interpreter/internal/ast/statements"
	"interpreter/internal/object"
)

func (v *ASTVisitor) VisitProgram(program *statements.Program) object.Object {
	return evalProgram(program.Statements, v)
}

func (v *ASTVisitor) VisitExpression(expression *statements.ExpressionStatement) object.Object {
	return EvaluateExpression(expression.Value, v)
}

func (v *ASTVisitor) VisitBlockStatement(block *statements.BlockStatement) object.Object {
	return evalBlockStatements(block.Statements, v)
}

func (v *ASTVisitor) VisitReturn(r *statements.ReturnStatement) object.Object {
	value := EvaluateExpression(r.Value, v)
	if isRuntimeError(value) {
		return value
	}
	return &object.Return{Value: value}
}

func (v *ASTVisitor) VisitLetStatement(let *statements.LetStatement) object.Object {
	value := EvaluateExpression(let.Value, v)
	if isRuntimeError(value) {
		return value
	}

	v.env.Set(let.Identifier.Literal(), value)

	// TODO: продумать контекст (environment) в текущем блоке
	return object.NULL
}
