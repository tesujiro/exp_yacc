package vm

import (
	"fmt"
	"reflect"
)

type Env struct {
	env    map[string]interface{}
	val    map[string]reflect.Value
	parent *Env
}

// Global Scope
func NewEnv() *Env {
	return &Env{
		env: make(map[string]interface{}),
		//val:    make(map[string]reflect.Value),
		parent: nil,
	}
}

func (e *Env) NewEnv() *Env {
	return &Env{
		env:    make(map[string]interface{}),
		val:    make(map[string]reflect.Value),
		parent: e,
	}
}

func (e *Env) Destroy() {
	if e.parent == nil {
		return
	}
	e.parent = nil
	e.env = nil
	return
}

func (e *Env) Set(k string, v interface{}) error {
	if _, ok := e.env[k]; ok {
		e.env[k] = v
		return nil
	}
	if e.parent == nil {
		return fmt.Errorf("unknown symbol '%s'", k)
	}
	return e.parent.Set(k, v)
}

func (e *Env) Define(k string, v interface{}) error {
	e.env[k] = v
	return nil
}

/*
func (e *Env) DefineValue(k string, v reflect.Value) error {
	e.val[k] = v
	return nil
}
*/

func (e *Env) Get(k string) (interface{}, error) {
	//fmt.Printf("Get(%#v)\n", k)
	if v, ok := e.env[k]; ok {
		return v, nil
	}
	if e.parent == nil {
		return nil, fmt.Errorf("unknown symbol '%s'", k)
	}
	return e.parent.Get(k)
}

/*
func (e *Env) GetValue(k string) (reflect.Value, error) {
	//fmt.Printf("Get(%#v)\n", k)
	if v, ok := e.val[k]; ok {
		return v, nil
		//return reflect.ValueOf(v), nil
	}
	if e.parent == nil {
		return reflect.ValueOf(nil), fmt.Errorf("unknown symbol '%s'", k)
	}
	return e.parent.GetValue(k)
}
*/
