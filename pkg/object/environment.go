package object

// Environment menyimpan pasangan variabel dan nilainya di memori
type Environment struct {
	store map[string]Object
	outer *Environment
}

// NewEnvironment membuat environment root (global scope)
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// NewEnclosedEnvironment membuat inner scope (untuk fungsi lokal)
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// Get mencari variabel di scope aktif atau merambat ke scope luar (outer)
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set mendaftarkan atau memperbarui nilai variabel
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
