package object

import "fmt"

type Null struct{}

func (i *Null) Inspect() string {
	return "null"
}

func (i *Null) Type() ObjectType {
	return NULL
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() ObjectType {
	return INTEGER
}

type Boolean struct {
	Value bool
}

func (i *Boolean) Inspect() string {
	return fmt.Sprintf("%t", i.Value)
}

func (i *Boolean) Type() ObjectType {
	return BOOLEAN
}
