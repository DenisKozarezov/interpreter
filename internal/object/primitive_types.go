package object

import "fmt"

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
	return "ERROR: " + e.Message
}

func (e *Error) Type() ObjectType {
	return ERROR_TYPE
}
