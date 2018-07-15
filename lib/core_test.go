package lib

import (
	"fmt"
	"go/token"
	"reflect"
	"testing"

	"github.com/rogpeppe/godef/go/scanner"
	"github.com/tesujiro/exp_yacc/parser"
	"github.com/tesujiro/exp_yacc/vm"
)

func aaa() {
}

func TestImportGoFuncs(t *testing.T) {
	tests := []struct {
		script     string
		result     interface{}
		errMessage string
	}{
		{script: "ret,err=println(\"hello!\");ret", result: 7},
		//{script: "func (x){return func(y) {return x*10+y};}(10)()", errMessage: "function wants 1 arguments but received 0"},
	}
	for _, test := range tests {
		env := vm.NewEnv()
		env = Import(env)

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
		if actual, err := vm.Run(parseResult, env); err != nil {
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
