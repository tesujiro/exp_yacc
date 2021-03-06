package vm

import (
	"fmt"
	"go/scanner"
	"go/token"
	"reflect"
	"testing"

	"github.com/tesujiro/exp_yacc/parser"
)

func TestNumbers(t *testing.T) {
	tests := []struct {
		script     string
		result     interface{}
		errMessage string
	}{
		{script: "true", result: true},
		{script: "false", result: false},
		{script: "nil", result: nil},
		{script: "1", result: 1},
		//{script: "9223372036854775807", result: 9223372036854775807},
		{script: "1.1", result: float64(1.1)},
		{script: "123", result: 123},
		{script: "123.456", result: float64(123.456)},
		{script: "\"abc\"", result: "abc"},
		{script: "+1+4", result: 5},
		{script: "-1+3", result: 2},
		{script: "1+1", result: 2},
		{script: "1+1.1", result: float64(2.1)},
		{script: "1.1+4", result: float64(5.1)},
		{script: "1.1+1.1", result: float64(2.2)},
		{script: "3-1.1", result: float64(1.9)},
		{script: "2.2-1.1", result: float64(1.1)},
		{script: "3-1-1", result: 1},
		{script: "3-(1-1)", result: 3},
		{script: "3*5", result: 15},
		{script: "1.5*2", result: float64(3)},
		{script: "5*1.2", result: float64(6)},
		{script: "15/5", result: 3},
		{script: "16/5", result: 3},
		{script: "3/1.5", result: float64(2)},
		{script: "3/0", errMessage: "devision by zero"},
		{script: "15%5", result: 0},
		{script: "16%5", result: 1},
		{script: "15%4.1", result: 3},
		{script: "\"a b c\"+\" d e f\"", result: "a b c d e f"},
		{script: "\"a b c\"-\" d e f\"", result: float64(0)},
		{script: "15.2%7.1", result: 1},
		{script: "a=1;b=2;a+b", result: 3},
		{script: "a=1;b=0.1;c=15;(a+b)*c", result: float64(16.5)},
		{script: "a=1;b=0.1;c=15;(a+b)*c/0.5", result: float64(33)},
		// bool
		{script: "a=1;b=2;a+1==b", result: true},
		{script: "a=1;b=2;a+1!=b", result: false},
		{script: "a=1;b=2;a<b", result: true},
		{script: "a=1;b=1;a<=b", result: true},
		{script: "a=1;b=2;a>b", result: false},
		{script: "a=1;b=1;a>=b", result: true},
		{script: "a=1;b=0.1;a>b", result: true},
		{script: "a=1;a==1", result: true},
		{script: "a=1;b=2;a==1&&b==2", result: true},
		{script: "a=1;b=2;a==2&&b==2", result: false},
		{script: "a=1;b=2;a==1||b==2", result: true},
		{script: "a=1;b=2;a==2||b==2", result: true},
		{script: "12&&34", errMessage: "cannot convert to bool"},
		// composite expression
		{script: "a=10;a++", result: 11},
		{script: "a=1.9;a++", result: 2.9},
		{script: "10++", errMessage: "Invalid operation"},
		{script: "a=\"a\";a++", errMessage: "Invalid operation"},
		{script: "a=10;a--", result: 9},
		{script: "a=2.9;a--", result: 1.9},
		{script: "10--", errMessage: "Invalid operation"},
		{script: "a=\"a\";a--", errMessage: "Invalid operation"},
		{script: "a=10;a+=2", result: 12},
		{script: "a=10;a+=2.5", result: 12.5},
		{script: "a=10;a+=2;a", result: 12},
		{script: "a=10;a-=2;a", result: 8},
		{script: "a=10;a*=2;a", result: 20},
		{script: "a=10;a/=2;a", result: 5},
		// multi result
		{script: "a,b=1,2;a", result: 1},
		{script: "a,b=1,2;b", result: 2},
		{script: "a,b=1,2,3;b", result: 2},
		{script: "a,b,c=1,2;b", result: 2},
		// if statement
		{script: "a=1;if a==1 { a=2 }", result: 2},
		{script: "a=1;if a==1 { a=2 };a", result: 2},
		{script: "a=1;if a==1 { env_test=2 };env_test", errMessage: "unknown symbol 'env_test'"},
		{script: "a=2;if a==1 { a=2 } else { a=3;b=4 }", result: 4},
		{script: "a=1;b=1;if a==1 { b=2 };b", result: 2},
		{script: "a=1;if a==1 { a=11 } else if a==2 { a=12 } else if a==3 { a=13 };a", result: 11},
		{script: "a=2;if a==1 { a=11 } else if a==2 { a=12 } else if a==3 { a=13 };a", result: 12},
		{script: "a=3;if a==1 { a=11 } else if a==2 { a=12 } else if a==3 { a=13 };a", result: 13},
		{script: "a=4;if a==1 { a=11 } else if a==2 { a=12 } else if a==3 { a=13 } else { a=14 };a", result: 14},
		{script: "a=1;if a==1 { env_test=11 } else if a==2 { env_test=12 } else { env_test=13 };env_test", errMessage: "unknown symbol 'env_test'"},
		{script: "a=2;if a==1 { env_test=11 } else if a==2 { env_test=12 } else { env_test=13 };env_test", errMessage: "unknown symbol 'env_test'"},
		{script: "a=3;if a==1 { env_test=11 } else if a==2 { env_test=12 } else { env_test=13 };env_test", errMessage: "unknown symbol 'env_test'"},
		{script: "a=1;if a==1 { a=2\n a=3 }", result: 3},
		{script: "a=1;if a==1 { a=2\n a=3\n }", result: 3},
		{script: "a=1;if a==1 {\n a=2\n a=3\n }", result: 3},
		{script: "a=1;if a==1 \n{\n a=2\n a=3\n }", errMessage: "unexpected ';'"},
		// for statement
		{script: "a=0;for{ if a==10 { break }; a= a+1 };a", result: 10},
		{script: "a=0;b=0;for{ a=a+1;if a==10 { break }; if b==5 {continue};b= b+1 };b", result: 5},
		{script: "a=0;for a<=10 { a= a+1 };a", result: 11},
		{script: "a=0;for b { a= a+1 };a", errMessage: "unknown symbol 'b'"},
		{script: "a=0;for a { a= a+1 };a", errMessage: "for condition type int cannot convert to bool"},
		{script: "a=0;for{ a=a+1;if a==10 { return a };};a", result: 10},
		{script: "a=0;for{ a=10;return a };a", result: 10},
		//{script: "Fn=func(){ a=0; for a<=100 { if a==10 {return 123};a=a+1} }; Fn()", result: 3},
		//{script: "Fn=func(){a=1; for { if a==3 {return a;} ; a=a+1}; Fn()", result: 3},
		// array
		{script: "ar=[1,2,3];ar", result: []interface{}{1, 2, 3}},
		{script: "ar=[1,\"2\",3];ar", result: []interface{}{1, "2", 3}},
		{script: "ar=[1,2,3];ar[0]", result: 1},
		{script: "ar=[1,2,3];ar[-1]", errMessage: "index out of range"},
		{script: "ar=[1,2,3];ar[3]", errMessage: "index out of range"},
		{script: "ar=[1,2,3];ar[\"x\"]", errMessage: "index cannot convert to int"},
		{script: "ar=1;ar[1]", errMessage: "type int does not support index operation"},
		{script: "ar=[1,2,3];ar[1]=100;ar", result: []interface{}{1, 100, 3}},
		{script: "ar=[1,2,3];ar[3]=4;ar", result: []interface{}{1, 2, 3, 4}},
		{script: "ar=[1,2,3];ar[4]=4", errMessage: "index out of range"},
		{script: "[1,2,3]+[4,5,6]", result: []interface{}{1, 2, 3, 4, 5, 6}},
		{script: "[1,2,3]+4", result: []interface{}{1, 2, 3, 4}},
		{script: "[1,2,3]+\"4\"", result: []interface{}{1, 2, 3, "4"}},
		// map
		{script: "m={1:\"aaa\",2:\"bbb\",3:\"ccc\"};m", result: map[interface{}]interface{}{1: "aaa", 2: "bbb", 3: "ccc"}},
		{script: "m={1:\"aaa\",2:\"bbb\",3:\"ccc\"};m[3]", result: "ccc"},
		{script: "m={1:\"aaa\",2:\"bbb\",3:\"ccc\"};m[0]", result: nil},
		{script: "m={1:\"aaa\",2:\"bbb\",3:\"ccc\"};len(m)", result: 3},
		{script: "m={1:\"aaa\",2:\"bbb\",3:\"ccc\"};m[3]=\"xxx\";m", result: map[interface{}]interface{}{1: "aaa", 2: "bbb", 3: "xxx"}},
		{script: "m={1:\"aaa\",2:\"bbb\",3:\"ccc\"};m[4]=\"xxx\";m", result: map[interface{}]interface{}{1: "aaa", 2: "bbb", 3: "ccc", 4: "xxx"}},
		{script: "m={1:\"aaa\",2:\"bbb\",3:\"ccc\"};m[4]=\"xxx\";m[3]+m[4]", result: "cccxxx"},
	}
	for _, test := range tests {
		env := NewEnv()
		//fmt.Println("*************************\nTEST SCRIPT:", test.script)
		l := new(parser.Lexer)

		fset := token.NewFileSet()                              // positions are relative to fset
		file := fset.AddFile("", fset.Base(), len(test.script)) // register input "file"
		l.Init(file, []byte(test.script), nil /* no error handler */, scanner.ScanComments)

		if parseResult, err := parser.Parse(l); err != nil {
			if test.errMessage == "" || err.Error() != test.errMessage {
				t.Errorf("Run error:%#v want%#v - script:%v\n", err, test.errMessage, test.script)
			}
		} else {
			//fmt.Println("src:", test.script, "\tast:", *parseResult[0])
			if actual, err := Run(parseResult, env); err != nil {
				if test.errMessage == "" || err.Error() != test.errMessage {
					t.Errorf("Run error:%#v want%#v - script:%v\n", err, test.errMessage, test.script)
				}
			} else if !reflect.DeepEqual(actual, test.result) {
				//} else if actual != test.result {
				t.Errorf("got %v\nwant %#v - script: %v", actual, test.result, test.script)
				fmt.Println("\tType:\t", reflect.TypeOf(actual))
				fmt.Println("\tValue:\t", reflect.ValueOf(actual))
				continue
			}
		}
	}
}

func TestFuncCall(t *testing.T) {
	tests := []struct {
		script     string
		result     interface{}
		errMessage string
	}{
		// User Defined Func
		{script: "a=10;a(10)", errMessage: "cannot call type int"},
		{script: "func Fn(a){3;};Fn(10)", result: 3},
		{script: "func Fn(a){a*10;};Fn(10)", result: 100},
		{script: "func Fn(a){a*10;};b=Fn(10);b*2", result: 200},
		{script: "func Add(a1,a2){return a1+a2;};Add(1,5)", result: 6},
		{script: "func Add(a1,a2){return a1+a2;};Add(1.1,1.2)", result: float64(2.3)},
		{script: "func Add(a1,a2){return a1+a2;};Add(\"hello,\",\"world!\")", result: "hello,world!"},
		{script: "func two(){1; return 2;3;};x=two();x", result: 2},
		{script: "Hundred=func Fn(){return 100;};a=Hundred();a", result: 100},
		// multi result function
		{script: "func Cross(a1,a2){return a2,a1;};Cross(1,5)", result: []interface{}{5, 1}},
		{script: "func Cross(a1,a2){return a2,a1;};x,y=Cross(1,5);x", result: 5},
		{script: "func Cross(a1,a2){return a2,a1;};x,y=Cross(1,5);y", result: 1},
		{script: "func Cross(a1,a2){return a2,a1;};Cross(\"a\",\"b\")", result: []interface{}{"b", "a"}},
		{script: "a=1;func Fn(){a=100;};Fn();a", result: 100},
		// anonymous func
		{script: "func (x){return x+100;}(10)", result: 110},
		{script: "func (x){return x+100;}()", errMessage: "function wants 1 arguments but received 0"},
		{script: "(1+1)(10)", errMessage: "cannot call type int"},
		{script: "Fn=func (x){return func(y) {return x*10+y};};Fn2=Fn(10);Fn2(2)", result: 102},
		// recursive call
		{script: "func Factorial(x){if x==1 {1} else { x*Factorial(x-1)}};Factorial(3)", result: 6},
		{script: "func Factorial(x){if x==1 {return 1} else { return x*Factorial(x-1)}};Factorial(3)", result: 6},
		// higher order function
		{script: "func (x){return func(y) {return x*10+y};}()(2)", errMessage: "function wants 1 arguments but received 0"},
		{script: "func (x){return func(y) {return x*10+y};}(10)()", errMessage: "function wants 1 arguments but received 0"},
		{script: "func (x){return func(y) {return x*10+y};}(10)(2)", result: 102},
		{script: "func Fibo(){x,y=0,1;return func(){x,y=y,x+y;return y}};f=Fibo();f();f();f();f();", result: 5},
		// higher order & recursive
		{script: "func mod(x){return func f(y){ if y<x {return y} else { return f(y-x) }}};mod3=mod(3);mod3(11);", result: 2},
		{script: "func f(x){if x==1 {return 1} else {return x*f(x-1)}};f(1)", result: 1},
		{script: "func f(x){if x==1 {return 1} else {return x*f(x-1)}};f(3)", result: 6},
		{script: "func f(x){if x==1 {return func(){ return 1}} else { return func(){ return x*f(x-1)() }}};f(1)()", result: 1},
		{script: "func f(x){if x==1 {return func(){ return 1}} else { return func(){ return x*f(x-1)() }}};f(3)()", result: 6}, //ERROR
		{script: "func f(x){if x==1 {return func(){ return x}} else { return func(){ return x*(x-1) }}};f(1)()", result: 1},
		{script: "func f(x){return func(){ return x}};f(1)()", result: 1},
		//PROBLEM
		{script: "func(x){return func(){return x}}(1)()", result: 1},
		{script: "func(x){if true {return func(){return x}}}(1)()", result: 1},

		// error while execute function
		{script: "func (x){return x+z}(10)", errMessage: "unknown symbol 'z'"},
		{script: "func (x){return func(y) {return x*10+y+z};}(10)(2)", errMessage: "unknown symbol 'z'"},
		// len opeartion
		{script: "len(\"1234567890\")", result: 10},
		{script: "len([1,2,3])", result: 3},
		{script: "len(1)", errMessage: "type int does not support len operation"},
	}
	for _, test := range tests {
		env := NewEnv()
		env.Define("println", reflect.ValueOf(fmt.Println))
		env.Define("printf", reflect.ValueOf(fmt.Printf))

		//fmt.Println("*************************\nTEST SCRIPT:", test.script)
		l := new(parser.Lexer)

		fset := token.NewFileSet()                              // positions are relative to fset
		file := fset.AddFile("", fset.Base(), len(test.script)) // register input "file"
		l.Init(file, []byte(test.script), nil /* no error handler */, scanner.ScanComments)

		parseResult, err := parser.Parse(l)
		if err != nil {
			if test.errMessage == "" || err.Error() != test.errMessage {
				t.Errorf("got error message:%#v want%#v - script:%v\n", err.Error(), test.errMessage, test.script)
			}
			continue
		}
		if actual, err := Run(parseResult, env); err != nil {
			if test.errMessage == "" || err.Error() != test.errMessage {
				t.Errorf("got error message:%#v want%#v - script:%v\n", err.Error(), test.errMessage, test.script)
			}
			//} else if actual != test.result {
			//} else if actual.(reflect.Value).Interface() != test.result.(reflect.Value).Interface() {
		} else if !reflect.DeepEqual(actual, test.result) {
			t.Errorf("got %#v\nwant %#v - script: %v", actual, test.result, test.script)
			fmt.Println("actual Type:", reflect.TypeOf(actual), "\tValue:", reflect.ValueOf(actual))
			fmt.Println("want Type:", reflect.TypeOf(test.result), "\tValue:", reflect.ValueOf(test.result))
			continue
		}
	}
}
