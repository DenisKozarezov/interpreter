package evaluator

import "interpreter/internal/object"

var builtins = map[string]*object.BuiltIn{
	"len": {
		Function: func(args ...object.Object) object.Object {
			argsLen := len(args)

			if argsLen != 1 {
				return newRuntimeError("wrong number of arguments. got = %d, want = 1", argsLen)
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newRuntimeError("argument to 'len' not supported, got %s", arg.Type())
			}
		},
	},
}
