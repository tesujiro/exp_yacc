package vm

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/tesujiro/exp_yacc/ast"
)

func defineFunc(funcExpr *ast.FuncExpr, env *Env) (interface{}, error) {
	// FuncType
	//argTypes := make([]reflect.Type, len(funcExpr.Args))
	argTypes := make([]reflect.Type, len(funcExpr.Args))
	for i := 0; i < len(argTypes); i++ {
		argTypes[i] = reflect.TypeOf(interface{}(int64(1)))
	}

	//TODO: variadic
	funcType := reflect.FuncOf(argTypes, []reflect.Type{reflect.TypeOf(interface{}(int64(1))), reflect.TypeOf(errors.New(""))}, false)

	// FuncDefinition
	//runVmFunction := func(in []interface{}) (interface{}, error) {
	runVmFunction := func(in []reflect.Value) []reflect.Value {
		newEnv := env.NewEnv()
		defer newEnv.Destroy()

		// set value to newEnv
		//
		//
		//

		if rv, err := Run(funcExpr.Stmts, newEnv); err != nil {
			return []reflect.Value{reflect.ValueOf(interface{}(nil)), reflect.ValueOf(err)}
		} else {
			return []reflect.Value{reflect.ValueOf(interface{}(rv)), reflect.ValueOf(error(nil))}
		}
		return []reflect.Value{reflect.ValueOf(interface{}(nil)), reflect.ValueOf(error(nil))}
	}

	fn := reflect.MakeFunc(funcType, runVmFunction)
	//fmt.Printf("fn.NumIn(): %#v\n", fn.Type().NumIn())
	//fmt.Printf("fn.Type(): %#v\n", fn.Type())
	//fmt.Printf("fn.Kind(): %#v\n", fn.Kind())
	//fmt.Printf("fn: %#v\n", fn)

	if funcExpr.Name != "" {
		if err := env.DefineValue(funcExpr.Name, fn); err != nil {
			//if err := env.Define(funcExpr.Name, fn); err != nil {
			return nil, err
		}
	}
	return fn, nil
}

func callFunc(callExpr *ast.CallExpr, env *Env) (interface{}, error) {
	fn, err := env.GetValue(callExpr.Name)
	//fn, err := env.Get(callExpr.Name)
	//fmt.Printf("fn: %#v\n", fn)
	//fmt.Printf("TypeOf(fn): %#v\n", reflect.TypeOf(fn))
	//fmt.Printf("ValueOf(fn): %#v\n", reflect.ValueOf(fn))
	//fmt.Printf("fn.Kind(): %#v\n", fn.Type().Kind())
	//fmt.Printf("fn.Call: %#v\n", reflect.ValueOf(fn).Call([]reflect.Value{}))
	if err != nil {
		return nil, err
	}
	//fmt.Printf("callExpr=%#v\tfn=%#v\n", callExpr, fn)
	//f := reflect.ValueOf(fn)
	f := fn
	if f.Kind() == reflect.Interface && !f.IsNil() {
		f = f.Elem()
	}
	//fmt.Printf("f.Kind()=%#v\n", f.Kind())
	//fmt.Printf("fn.Call(aaa)=%v\n", f.Call(args))
	if args, err := callArgs(f, callExpr, env); err != nil {
		return nil, err
	} else {
		fmt.Printf("args: %#v\n", args)
		refvals := f.Call(args)
		return refvals, nil
	}
	return nil, nil
}

func callArgs(f reflect.Value, callExpr *ast.CallExpr, env *Env) ([]reflect.Value, error) {
	//if f.Type().IsVariadic() {
	//return nil, errors.New("sorry! TODO:Variadic")
	//}
	//fmt.Printf("function args should be:%d received:%d\n", f.Type().NumIn(), len(callExpr.SubExprs))
	//fmt.Printf("f.Type():%v\n", f.Type())
	//fmt.Printf("f.Kind(): %#v\n", f.Kind())
	if f.Type().NumIn() < 1 {
		return []reflect.Value{}, nil
	}
	// TODO: check args length
	//
	args := make([]reflect.Value, 0, f.Type().NumIn())
	fmt.Printf("NumIn: %#v\n", f.Type().NumIn())
	for _, subExpr := range callExpr.SubExprs {
		if arg, err := evalExpr(subExpr, env); err != nil {
			return nil, err
		} else {
			args = append(args, reflect.ValueOf(arg))
			//fmt.Printf("arg: %#v\n", arg)
		}
	}
	return args, nil
}
