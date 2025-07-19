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
