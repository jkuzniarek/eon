package card


type Env struct {
	store map[string]Card
}


func NewEnv() *Env {
	s := make(map[string]Card)
	return &Env{store: s}
}

func (e *Env) Get(name string) (Card, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

func (e *Env) Set(name string, val Card) Card {
	e.store[name] = val
	return val
}