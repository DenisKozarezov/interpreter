package object

import (
	"bytes"
	"fmt"
	"strings"
)

// Below are defined the so-called `singleton values`.
//
// These are objects that are vitally important in order to avoid allocating the same
// type of entity with the same value every time. For example, there is a boolean
// variable that can only have two values: true and false. So we don't really need to
// constantly allocate the same object from heap during the runtime phase. Thus, for
// optimization purposes, it is better to point to pre-allocated objects that have
// already been created with the value of true/false, etc...
var (
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
	NULL  = &Null{}
)

type Null struct{}

func (i *Null) Inspect() string {
	return "null"
}

func (i *Null) Type() ObjectType {
	return NULL_TYPE
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() ObjectType {
	return INTEGER_TYPE
}

type Boolean struct {
	Value bool
}

func (i *Boolean) Inspect() string {
	return fmt.Sprintf("%t", i.Value)
}

func (i *Boolean) Type() ObjectType {
	return BOOLEAN_TYPE
}

func NativeBooleanToObject(input bool) Object {
	if input {
		return TRUE
	}
	return FALSE
}

func ObjectToNativeBoolean(obj Object) bool {
	switch obj {
	case NULL, FALSE:
		return false
	case TRUE:
		return true
	default:
		return true
	}
}

type Return struct {
	Value Object
}

func (r *Return) Inspect() string {
	return r.Value.Inspect()
}

func (r *Return) Type() ObjectType {
	return RETURN_TYPE
}

type Error struct {
	Message string
}

func (e *Error) Inspect() string {
	return "runtime error: " + e.Message
}

func (e *Error) Type() ObjectType {
	return ERROR_TYPE
}

type node interface {
	fmt.Stringer
	Literal() string
}

type Function struct {
	Args        []fmt.Stringer
	Body        node
	Environment *Environment
}

func (f *Function) Inspect() string {
	var out bytes.Buffer
	var params []string
	for i := range f.Args {
		params = append(params, f.Args[i].String())
	}
	out.WriteString("fn (")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

func (f *Function) Type() ObjectType {
	return FUNCTION_TYPE
}

type String struct {
	Value string
}

func (s *String) Inspect() string {
	return s.Value
}

func (s *String) Type() ObjectType {
	return STRING_TYPE
}

type BuiltinFunction = func(...Object) Object

type BuiltIn struct {
	Function BuiltinFunction
}

func (b *BuiltIn) Inspect() string {
	return "built-in function"
}

func (b *BuiltIn) Type() ObjectType {
	return BUILTIN_TYPE
}

type Array struct {
	Items []Object
}

func (a *Array) Inspect() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")

	items := make([]string, len(a.Items))
	for i := 0; i < len(a.Items); i++ {
		items[i] = a.Items[i].Inspect()
	}
	buffer.WriteString(strings.Join(items, ", "))
	buffer.WriteString("]")
	return buffer.String()
}

func (a *Array) Type() ObjectType {
	return ARRAY_TYPE
}
