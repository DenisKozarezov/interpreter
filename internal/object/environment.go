package object

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnvironment() *Environment {
	return &Environment{store: make(map[string]Object)}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func (e *Environment) Set(identifier string, val Object) Object {
	e.store[identifier] = val
	return val
}

func (e *Environment) Get(identifier string) (Object, bool) {
	obj, ok := e.store[identifier]
	if !ok && e.outer != nil {
		return e.outer.Get(identifier)
	}
	return obj, ok
}
