package evaluator

import (
	"fmt"
	"interpreter/internal/object"
)

func newRuntimeError(format string, args ...any) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, args...)}
}

func isRuntimeError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_TYPE
	}
	return false
}
