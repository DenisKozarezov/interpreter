package object

type Environment struct {
	store map[string]Object
}

func NewEnvironment() *Environment {
	return &Environment{store: make(map[string]Object)}
}

func (e *Environment) Set(identifier string, val Object) Object {
	e.store[identifier] = val
	return val
}

func (e *Environment) Get(identifier string) (Object, bool) {
	obj, ok := e.store[identifier]
	return obj, ok
}
