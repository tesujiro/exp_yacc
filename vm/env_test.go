package vm

import (
	"fmt"
	"testing"
)

func TestEnv(t *testing.T) {
	cases := []struct {
		k string
		v interface{}
	}{
		{k: "abc", v: 1},
	}
	for _, c := range cases {
		e := NewEnv()
		if err := e.Define(c.k, c.v); err != nil {
			t.Errorf("Env.Set error :%v", err)
			continue
		} else if actual, err := e.Get(c.k); err != nil {
			t.Errorf("Env.Get error :%v", err)
			continue
		} else if actual != c.v {
			t.Errorf("got %v\nwant %v", actual, c.v)
			continue
		}
		fmt.Printf("env=%#v\n", e)
	}
}

func TestChildEnv(t *testing.T) {
	tests := []struct {
		k string
		v interface{}
	}{
		{k: "int", v: 123},
		{k: "float", v: 1.1},
	}

	root := NewEnv()
	for _, test := range tests {
		if err := root.Define(test.k, test.v); err != nil {
			t.Errorf("Env.Set error :%v", err)
			return
		} else if actual, err := root.Get(test.k); err != nil {
			t.Errorf("Env.Get error :%v", err)
			return
		} else if actual != test.v {
			t.Errorf("got %v\nwant %v", actual, test.v)
			return
		}
	}
	fmt.Printf("root=%#v\n", root)

	child := root.NewEnv()
	for _, test := range tests {
		if err := child.Define(test.k, test.v); err != nil {
			t.Errorf("Env.Set error :%v", err)
			return
		} else if actual, err := child.Get(test.k); err != nil {
			t.Errorf("Env.Get error :%v", err)
			return
		} else if actual != test.v {
			t.Errorf("got %v\nwant %v", actual, test.v)
			return
		}
	}
	fmt.Printf("child=%#v\n", child)

	child.Destroy()
	if child.parent != nil {
		t.Errorf("child.parent != nil")
	}
	fmt.Printf("after Destroy child=%#v\n", child)
	fmt.Printf("root=%#v\n", root)

}
