package vm

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/tesujiro/exp_yacc/parser"
)

func main() {
	fmt.Println("vim-go")
}

func TestNumbers(t *testing.T) {
	tests := []struct {
		script     string
		result     interface{}
		errMessage string
	}{
		{script: "true", result: true},
		{script: "false", result: false},
		{script: "nil", result: nil},
		{script: "1", result: int64(1)},
		{script: "1.1", result: float64(1.1)},
		{script: "123", result: int64(123)},
		{script: "123.456", result: float64(123.456)},
		{script: "\"abc\"", result: "abc"},
		{script: "+1+4", result: int64(5)},
		{script: "-1+3", result: int64(2)},
		{script: "1+1", result: int64(2)},
		{script: "1+1.1", result: float64(2.1)},
		{script: "1.1+4", result: float64(5.1)},
		{script: "1.1+1.1", result: float64(2.2)},
		{script: "3-1.1", result: float64(1.9)},
		{script: "2.2-1.1", result: float64(1.1)},
		{script: "3-1-1", result: int64(1)},
		{script: "3-(1-1)", result: int64(3)},
		{script: "3*5", result: int64(15)},
		{script: "1.5*2", result: float64(3)},
		{script: "5*1.2", result: float64(6)},
		{script: "15/5", result: int64(3)},
		{script: "16/5", result: int64(3)},
		{script: "3/1.5", result: float64(2)},
		{script: "3/0", errMessage: "devision by zero"},
		{script: "15%5", result: int64(0)},
		{script: "16%5", result: int64(1)},
		{script: "15%4.1", result: int64(3)},
		{script: "\"a b c\"+\" d e f\"", result: "a b c d e f"},
		{script: "\"a b c\"-\" d e f\"", result: float64(0)},
		{script: "15.2%7.1", result: int64(1)},
		{script: "a=1;b=2;a+b", result: int64(3)},
		{script: "a=1;b=2;a+1==b", result: true},
		{script: "a=1;b=2;a+1!=b", result: false},
		{script: "a=1;b=2;a<b", result: true},
		{script: "a=1;b=1;a<=b", result: true},
		{script: "a=1;b=2;a>b", result: false},
		{script: "a=1;b=1;a>=b", result: true},
		{script: "a=1;b=0.1;a>b", result: true},
		{script: "a=1;b=0.1;c=15;(a+b)*c", result: float64(16.5)},
		{script: "a=1;b=0.1;c=15;(a+b)*c/0.5", result: float64(33)},
		// multi result
		{script: "a,b=1,2;a", result: int64(1)},
		{script: "a,b=1,2;b", result: int64(2)},
		// if
		{script: "a=1;if a==1 { a=2 }", result: int64(2)},
		{script: "a=1;if a==1 { a=2 };a", result: int64(2)},
		{script: "a=1;if a==1 { env_test=2 };env_test", errMessage: "unknown symbol 'env_test'"},
		{script: "a=2;if a==1 { a=2 } else { a=3;b=4 }", result: int64(4)},
		{script: "a=1;b=1;if a==1 { b=2 };b", result: int64(2)},
		{script: "a=1;if a==1 { a=11 } else if a==2 { a=12 } else if a==3 { a=13 };a", result: int64(11)},
		{script: "a=2;if a==1 { a=11 } else if a==2 { a=12 } else if a==3 { a=13 };a", result: int64(12)},
		{script: "a=3;if a==1 { a=11 } else if a==2 { a=12 } else if a==3 { a=13 };a", result: int64(13)},
		{script: "a=4;if a==1 { a=11 } else if a==2 { a=12 } else if a==3 { a=13 } else { a=14 };a", result: int64(14)},
		{script: "a=1;if a==1 { env_test=11 } else if a==2 { env_test=12 } else { env_test=13 };env_test", errMessage: "unknown symbol 'env_test'"},
		{script: "a=2;if a==1 { env_test=11 } else if a==2 { env_test=12 } else { env_test=13 };env_test", errMessage: "unknown symbol 'env_test'"},
		{script: "a=3;if a==1 { env_test=11 } else if a==2 { env_test=12 } else { env_test=13 };env_test", errMessage: "unknown symbol 'env_test'"},
		{script: "a=1;if a==1 { a=2\n a=3 }", result: int64(3)},
		{script: "a=1;if a==1 { a=2\n a=3\n }", result: int64(3)},
	}
	for _, test := range tests {
		env := NewEnv()
		l := new(parser.Lexer)
		l.Init(strings.NewReader(test.script))
		parseResult, _ := parser.Parse(l)
		if actual, err := Run(parseResult, env); err != nil {
			if test.errMessage == "" || err.Error() != test.errMessage {
				t.Errorf("Run error:%#v want%#v - script:%v\n", err, test.errMessage, test.script)
			}
		} else if actual != test.result {
			t.Errorf("got %#v\nwant %#v - script: %v", actual, test.result, test.script)
			continue
		}
	}
}
func TestFuncCall(t *testing.T) {
	tests := []struct {
		script     string
		result     interface{}
		errMessage string
	}{
		{script: "ret,err=println(\"hello!\");ret", result: 7},
		//{script: "println(\"hello!\")", result: []reflect.Value{reflect.ValueOf(7), reflect.ValueOf(error(nil))}},
		{script: "a=10;a(10)", errMessage: "cannot call type int64"},
		{script: "func Fn(a){3;};Fn(10)", result: int64(3)},
		{script: "func Fn(a){a*10;};Fn(10)", result: int64(100)},
		{script: "func Fn(a){a*10;};b=Fn(10);b*2", result: int64(200)},
		{script: "func Add(a1,a2){return a1+a2;};Add(1,5)", result: int64(6)},
		{script: "func Add(a1,a2){return a1+a2;};Add(1.1,1.2)", result: float64(2.3)},
		{script: "func Add(a1,a2){return a1+a2;};Add(\"hello,\",\"world!\")", result: "hello,world!"},
		{script: "func two(){1; return 2;3;};x=two();x", result: int64(2)},
		// multi result
		{script: "func Cross(a1,a2){return a2,a1;};Cross(1,5)", result: []interface{}{int64(5), int64(1)}},
		{script: "func Cross(a1,a2){return a2,a1;};x,y=Cross(1,5);x", result: int64(5)},
		{script: "func Cross(a1,a2){return a2,a1;};x,y=Cross(1,5);y", result: int64(1)},
		{script: "func Cross(a1,a2){return a2,a1;};Cross(\"a\",\"b\")", result: []interface{}{"b", "a"}},
		{script: "a=1;func Fn(){a=100;};Fn();a", result: int64(100)},
	}
	for _, test := range tests {
		env := NewEnv()
		env.Define("println", reflect.ValueOf(fmt.Println))
		env.Define("printf", reflect.ValueOf(fmt.Printf))

		l := new(parser.Lexer)
		l.Init(strings.NewReader(test.script))
		parseResult, err := parser.Parse(l)
		if err != nil {
			fmt.Printf("Syntax error: %v \n", err)
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
