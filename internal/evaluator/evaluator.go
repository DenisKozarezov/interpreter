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
