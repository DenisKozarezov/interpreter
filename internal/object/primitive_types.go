package object

import (
	"bytes"
	"fmt"
	"strings"
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

// Ниже представлены так называемые значения-синглтоны (singleton values).
// Это объекты, которые необходимы для того, чтобы не аллоцировать каждый раз
// однотипные сущности с одним и тем же значением. Например, есть булева переменная,
// которая может иметь только два значения: true и false.
//
// В целях оптимизации лучше указывать на уже заранее созданные объекты со значением true/false.
var (
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
	NULL  = &Null{}
)

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
