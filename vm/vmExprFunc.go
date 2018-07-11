package vm

import (
	"errors"
	"fmt"
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

func callAnonymousFunc(anonymousCallExpr *ast.AnonymousCallExpr, env *Env) (interface{}, error) {
	ace := anonymousCallExpr
	debug.Printf("anonCallExpr:%#v\n", ace)
	if result, err := evalExpr(ace.Expr, env); err != nil {
		return nil, err
	} else {
		if rv, ok := result.(reflect.Value); !ok {
			return nil, errors.New("cannot call type " + reflect.TypeOf(result).String())
		} else {
			if rv.Type().Kind() != reflect.Func {
				return nil, errors.New("cannot call type " + reflect.TypeOf(result).String())
			}
			return callFunc(&ast.CallExpr{Func: rv, SubExprs: ace.SubExprs}, env)
		}
	}
}

func callFunc(callExpr *ast.CallExpr, env *Env) (interface{}, error) {
	var f reflect.Value
	if callExpr.Name != "" {
		fn, err := env.Get(callExpr.Name) // fn: interface{}
		if err != nil {
			return nil, err
		}
		var ok bool
		f, ok = fn.(reflect.Value) // interface{} ==> reflect.Value
		if !ok {
			return nil, errors.New("cannot call type " + reflect.TypeOf(fn).String())
		}
	} else {
		f = callExpr.Func
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
		/*
			fmt.Printf("=====>args: %v\n", args)
			for i, arg := range args {
				fmt.Printf("=====>===>arg[%d]: %v\n", i, arg)
			}
		*/

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
	if f.Type().NumIn() < 1 {
		return []reflect.Value{}, nil
	}
	if f.Type().NumIn() != len(callExpr.SubExprs) {
		return []reflect.Value{}, fmt.Errorf("function wants %v arguments but received %v", f.Type().NumIn(), len(callExpr.SubExprs))
	}
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
