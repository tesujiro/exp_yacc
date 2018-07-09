package vm

import (
	"errors"
	"reflect"

	"github.com/tesujiro/exp_yacc/ast"
	"github.com/tesujiro/exp_yacc/debug"
)

func defineFunc(funcExpr *ast.FuncExpr, env *Env) (interface{}, error) {
	// FuncType
	inType := make([]reflect.Type, len(funcExpr.Args))
	for i := 0; i < len(inType); i++ {
		inType[i] = reflect.TypeOf(reflect.Value{})
	}

	outType := []reflect.Type{reflect.TypeOf(reflect.Value{}), reflect.TypeOf(reflect.Value{})}
	isVariadic := false //TODO: variadic
	funcType := reflect.FuncOf(inType, outType, isVariadic)

	// FuncDefinition
	//runVmFunction := func(in []interface{}) (interface{}, error) {
	runVmFunction := func(in []reflect.Value) []reflect.Value {
		newEnv := env.NewEnv()
		defer newEnv.Destroy()

		for i, arg := range funcExpr.Args {
			//val := reflect.ValueOf(in[i]).Interface()
			//val := reflect.ValueOf(in[i]).Interface().(reflect.Value).Interface()
			val := reflect.ValueOf(in[i]).Interface().(reflect.Value).Interface().(reflect.Value).Interface()
			debug.Printf("arg[%v]: %#v\tType:%v\tValue:%v\n", i, in[i], reflect.TypeOf(val), reflect.ValueOf(val))
			if err := newEnv.Set(arg, val); err != nil {
				if err := newEnv.Define(arg, val); err != nil {
					return []reflect.Value{reflect.ValueOf(interface{}(nil)), reflect.ValueOf(err)}
				}
			}
		}
		debug.Printf("newEnv: %#v\n", *newEnv)

		var errValue reflect.Value
		var errorType = reflect.ValueOf([]error{nil}).Index(0).Type()
		var reflectValueErrorNilValue = reflect.ValueOf(reflect.New(errorType).Elem())
		errValue = reflectValueErrorNilValue

		if rv, err := Run(funcExpr.Stmts, newEnv); err != nil {
			// TODO: maybe buggy
			return []reflect.Value{reflect.ValueOf(interface{}(nil)), reflect.ValueOf(err)} //TODO
		} else {
			return []reflect.Value{reflect.ValueOf(reflect.ValueOf(rv)), reflect.ValueOf(errValue)}
		}
	}

	fn := reflect.MakeFunc(funcType, runVmFunction)

	if funcExpr.Name != "" {
		if err := env.Define(funcExpr.Name, fn); err != nil {
			return nil, err
		}
	}
	return fn, nil
}

func callFunc(callExpr *ast.CallExpr, env *Env) (interface{}, error) {
	fn, err := env.Get(callExpr.Name) // fn: interface{}
	if err != nil {
		return nil, err
	}
	f, ok := fn.(reflect.Value) // interface{} ==> reflect.Value
	if !ok {
		return nil, errors.New("cannot call type " + reflect.TypeOf(fn).String())
	}
	debug.Println("func kind:", f.Kind())
	if f.Kind() != reflect.Func {
		return nil, errors.New("cannot call type " + f.Type().String())
	}
	/*
		if f.Kind() == reflect.Interface && !f.IsNil() {
			f = f.Elem()
		}
	*/
	if args, err := callArgs(f, callExpr, env); err != nil {
		return nil, err
	} else {
		debug.Printf("args: %#v\n", args)

		// Call Function
		refvals := f.Call(args)

		debug.Println("refvals[0]: type ", reflect.TypeOf(refvals[0]), "\tValue ", reflect.ValueOf(refvals[0]))
		debug.Println("refvals[0]: type ", reflect.TypeOf(refvals[0].Interface()), "\tValue ", reflect.ValueOf(refvals[0].Interface()))
		if isGoFunc(f.Type()) {
			// Golang Pacakage Funcion
			result := make([]interface{}, len(refvals))
			for k, v := range refvals {
				result[k] = v.Interface()
			}
			return result, nil
		} else {
			// User Defined Funcion
			result := refvals[0].Interface().(reflect.Value).Interface()
			return result, nil
		}
	}
	return nil, nil
}

func isGoFunc(rt reflect.Type) bool {
	if rt.NumOut() != 2 || rt.Out(0) != reflect.TypeOf(reflect.Value{}) || rt.Out(1) != reflect.TypeOf(reflect.Value{}) {
		return true
	}
	return false
}

func callArgs(f reflect.Value, callExpr *ast.CallExpr, env *Env) ([]reflect.Value, error) {
	//if f.Type().IsVariadic() {
	//return nil, errors.New("sorry! TODO:Variadic")
	//}
	//fmt.Printf("function args should be:%d received:%d\n", f.Type().NumIn(), len(callExpr.SubExprs))
	if f.Type().NumIn() < 1 {
		return []reflect.Value{}, nil
	}
	// TODO: check args length
	//
	args := make([]reflect.Value, f.Type().NumIn(), f.Type().NumIn())
	for k, subExpr := range callExpr.SubExprs {
		if arg, err := evalExpr(subExpr, env); err != nil {
			return nil, err
		} else {
			// User Defined Funcion
			args[k] = reflect.ValueOf(reflect.ValueOf(arg))
			// TODO: Golang Pacakage Funcion
			//
		}
	}
	return args, nil
}
