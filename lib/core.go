package lib

import (
	"fmt"
	"reflect"

	"github.com/tesujiro/exp_yacc/vm"
)

func Import(env *vm.Env) *vm.Env {
	env.Define("println", reflect.ValueOf(fmt.Println))
	env.Define("printf", reflect.ValueOf(fmt.Printf))

	return env
}
