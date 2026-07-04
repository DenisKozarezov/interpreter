package evaluator

import (
	"interpreter/internal/object"

	"log"
)

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
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Items))}
			default:
				return newRuntimeError("argument to 'len' not supported, got %s", arg.Type())
			}
		},
	},

	"log": {
		Function: func(args ...object.Object) object.Object {
			if len(args) == 0 {
				return newRuntimeError("no arguments were passed")
			}

			logArgs := make([]any, len(args))
			for i := range args {
				logArgs[i] = args[i].Inspect()
			}

			log.Println(logArgs...)
			return object.NULL
		},
	},

	"logf": {
		Function: func(args ...object.Object) object.Object {
			argsLen := len(args)

			switch argsLen {
			case 0:
				return newRuntimeError("no arguments were passed")
			case 1:
				log.Println(args[0].Inspect())
				return object.NULL
			default:
				format := args[0]
				if format.Type() != object.STRING_TYPE {
					return newRuntimeError("the first argument is invalid, expected STRING, but got = %s", format.Type())
				}

				logArgs := make([]any, 0, argsLen-1)
				for i := 1; i < argsLen; i++ {
					switch args[i].Type() {
					case object.INTEGER_TYPE:
						logArgs = append(logArgs, args[i].(*object.Integer).Value)
					case object.BOOLEAN_TYPE:
						logArgs = append(logArgs, args[i].(*object.Boolean).Value)
					default:
						logArgs = append(logArgs, args[i].Inspect())
					}
				}

				log.Printf(format.Inspect(), logArgs...)
				return object.NULL
			}
		},
	},
}
