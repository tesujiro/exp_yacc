package vm

import (
	"reflect"

	"github.com/tesujiro/exp_yacc/ast"
	"github.com/tesujiro/exp_yacc/debug"
)

func defineFunc(funcExpr *ast.FuncExpr, env *Env) (interface{}, error) {
	// FuncType
	//argTypes := make([]reflect.Type, len(funcExpr.Args))
	inType := make([]reflect.Type, len(funcExpr.Args))
	for i := 0; i < len(inType); i++ {
		//inType[i] = reflect.TypeOf(int64(1))
		inType[i] = reflect.TypeOf(int64(1))
		//inType[i] = reflect.TypeOf(interface{}(int64(1)))
	}

	//TODO: variadic
	outType := []reflect.Type{reflect.TypeOf(reflect.Value{}), reflect.TypeOf(reflect.Value{})}
	//outType := []reflect.Type{reflect.TypeOf(interface{}(int64(1))), reflect.TypeOf(reflect.Value{})}
	//outType := []reflect.Type{reflect.TypeOf(int64(1)), reflect.TypeOf(reflect.Value{})}
	funcType := reflect.FuncOf(inType, outType, false)
	//funcType := reflect.FuncOf(inType, []reflect.Type{reflect.TypeOf(interface{}(int64(1))), reflect.TypeOf(errors.New(""))}, false)

	// FuncDefinition
	//runVmFunction := func(in []interface{}) (interface{}, error) {
	runVmFunction := func(in []reflect.Value) []reflect.Value {
		newEnv := env.NewEnv()
		defer newEnv.Destroy()

		// set value to newEnv
		for i, arg := range funcExpr.Args {
			if err := newEnv.Set(arg, in[i]); err != nil {
				if err := newEnv.Define(arg, reflect.ValueOf(in[i]).Interface().(reflect.Value).Interface()); err != nil {
					return []reflect.Value{reflect.ValueOf(interface{}(nil)), reflect.ValueOf(err)}
				}
			}
		}
		debug.Printf("newEnv: %#v\n", *newEnv)
		//
		var errValue reflect.Value
		var errorType = reflect.ValueOf([]error{nil}).Index(0).Type()
		var reflectValueErrorNilValue = reflect.ValueOf(reflect.New(errorType).Elem())
		errValue = reflectValueErrorNilValue

		if rv, err := Run(funcExpr.Stmts, newEnv); err != nil {
			return []reflect.Value{reflect.ValueOf(interface{}(nil)), reflect.ValueOf(err)}
		} else {
			//fmt.Println("rv: Type", reflect.TypeOf(rv), "\tValue:", reflect.ValueOf(rv))
			//return []reflect.Value{reflect.ValueOf(rv), reflect.ValueOf(errValue)}
			return []reflect.Value{reflect.ValueOf(reflect.ValueOf(rv)), reflect.ValueOf(errValue)}
			//return []reflect.Value{reflect.ValueOf(rv), reflect.ValueOf(error(nil))}
			//return []reflect.Value{reflect.ValueOf(interface{}(rv)), reflect.ValueOf(error(nil))}
			//return []reflect.Value{reflect.ValueOf(reflect.ValueOf(interface{}(rv))), reflect.ValueOf(error(nil))}
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
	f, err := env.GetValue(callExpr.Name)
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
	if f.Kind() == reflect.Interface && !f.IsNil() {
		f = f.Elem()
	}
	//fmt.Printf("f.Kind()=%#v\n", f.Kind())
	//fmt.Printf("fn.Call(aaa)=%v\n", f.Call(args))
	if args, err := callArgs(f, callExpr, env); err != nil {
		return nil, err
	} else {
		debug.Printf("args: %#v\n", args)
		refvals := f.Call(args)
		debug.Println("refvals: type ", reflect.TypeOf(refvals), "\tValue ", reflect.ValueOf(refvals))
		debug.Println("refvals[0]: type ", reflect.TypeOf(refvals[0]), "\tValue ", reflect.ValueOf(refvals[0]))
		debug.Println("refvals[0]: type ", reflect.TypeOf(refvals[0].Interface()), "\tValue ", reflect.ValueOf(refvals[0].Interface()))
		//return reflect.ValueOf(refvals[0]), nil
		return refvals[0].Interface().(reflect.Value).Interface(), nil
		//return refvals[0].Interface(), refvals[1].Interface()
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
	//args := make([]reflect.Value, f.Type().NumIn())
	args := make([]reflect.Value, f.Type().NumIn(), f.Type().NumIn())
	//args := []reflect.Value{}
	//fmt.Printf("NumIn: %#v\n", f.Type().NumIn())
	//fmt.Printf("len(SubExprs): %#v\n", len(callExpr.SubExprs))
	for k, subExpr := range callExpr.SubExprs {
		if arg, err := evalExpr(subExpr, env); err != nil {
			return nil, err
		} else {
			//args[k] = reflect.ValueOf(reflect.ValueOf(arg))
			args[k] = reflect.ValueOf(arg)
			//args = append(args, arg)
			//args = append(args, reflect.ValueOf(arg))
			//args = append(args, reflect.ValueOf(reflect.ValueOf(arg)))
			//args = append(args, reflect.ValueOf(reflect.ValueOf(reflect.ValueOf(arg))))
			//fmt.Printf("arg: %#v\n", arg)
		}
	}
	return args, nil
}
