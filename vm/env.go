package vm

type Env struct {
	env    map[string]interface{}
	parent *Env
}

// Global Scope
func NewEnv() *Env {
	return &Env{
		env:    make(map[string]interface{}),
		parent: nil,
	}
}

func (e *Env) NewEnv() *Env {
	return &Env{
		env:    make(map[string]interface{}),
		parent: e,
	}
}

func (e *Env) Set(k string, v interface{}) error {
	e.env[k] = v
	return nil
}

func (e *Env) Get(k string) (interface{}, error) {
	return e.env[k], nil
}
