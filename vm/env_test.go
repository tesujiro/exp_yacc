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
		if err := e.Set(c.k, c.v); err != nil {
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
