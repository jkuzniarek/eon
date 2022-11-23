package card


type Env struct {
	store map[string]Card
	parent *Env
}


func NewEnv() *Env {
	s := make(map[string]Card)
	return &Env{store: s, parent: nil}
}

func NewChildEnv(parent *Env) *Env {
	env := NewEnv()
	env.parent = parent
	return env
}

func (e *Env) Get(name string) (Card, bool) {
	obj, ok := e.store[name]
	if !ok && e.parent != nil {
		obj, ok = e.parent.Get(name)
	}
	return obj, ok
}

func (e *Env) Set(name string, val Card) Card {
	e.store[name] = val
	return val
}