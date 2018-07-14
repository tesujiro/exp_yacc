package main

import (
	"bufio"
	"flag"
	"fmt"
	"go/token"
	"io/ioutil"
	"os"

	"github.com/rogpeppe/godef/go/scanner"
	"github.com/tesujiro/exp_yacc/debug"
	"github.com/tesujiro/exp_yacc/parser"
	"github.com/tesujiro/exp_yacc/vm"
)

var dbg = flag.Bool("d", false, "debug option")
var file string

func main() {
	flag.Parse()
	file = flag.Arg(0)

	if *dbg {
		debug.On()
	}

	if file != "" {
		runScriptFile(file)
	} else {
		run()
	}
}

func runScriptFile(file string) {
	sourceBytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("ReadFile error:", err)
		return
	}
	source := string(sourceBytes)

	env := vm.NewEnv()

	l := new(parser.Lexer)
	//l.Init(strings.NewReader(source))

	fset := token.NewFileSet()                      // positions are relative to fset
	f := fset.AddFile("", fset.Base(), len(source)) // register input "file"
	l.Init(f, []byte(source), nil /* no error handler */, scanner.ScanComments)

	ast, parseError := parser.Parse(l)
	debug.Printf("%#v\n", ast)
	if parseError != nil {
		fmt.Printf("Syntax error: %v \n", parseError)
		//fmt.Printf("Syntax error: %v at %v\n", e, l.Position) //TODO
		return
	}
	if res, err := vm.Run(ast, env); err != nil {
		fmt.Printf("Eval error:%v\n", err)
		return
	} else {
		debug.Printf("ENV=%#v\n", env)
		fmt.Printf("%#v\n", res)
	}
	return
}

/*
func dump(stmts []ast.Stmt) {
	var dump_helper func([]interface{}, string)
	dump_helper = func(stmts []interface{}, indent string) {
		for _, stmt := range stmts {
			fmt.Println(indent, stmt)
			if reflect.TypeOf(stmt).Kind() == reflect.Slice {
				dump_helper(reflect.ValueOf(stmt).Elem(), indent+"\t")
			}
		}
	}
	dump_helper([]interface{}(stmts), "")
}
*/

func run() {
	env := vm.NewEnv()
	line_scanner := bufio.NewScanner(os.Stdin) // This Scanner
	var source string

	for line_scanner.Scan() {
		source += line_scanner.Text()
		//if source == "" {
		//continue
		//}
		if source == "exit" || source == "quit" {
			break
		}

		l := new(parser.Lexer)
		//l.Init(strings.NewReader(source))

		fset := token.NewFileSet()                         // positions are relative to fset
		file := fset.AddFile("", fset.Base(), len(source)) // register input "file"
		l.Init(file, []byte(source), nil /* no error handler */, scanner.ScanComments)

		ast, parseError := parser.Parse(l)
		for _, stmt := range ast {
			debug.Printf("%#v\n", stmt)
		}
		if parseError != nil {
			debug.Println("[", parseError.Error(), "]")
			//if parseError.Error() == "unexpected $end" || parseError.Error() == "comment not terminated" { //Does not work
			if parseError.Error() == "unexpected $end" {
				// note: scanner.Scan() does not return "end of line" ,
				// this is just for separating tokens
				source += "\n"
				//fmt.Println("source;[" + source + "]")
				continue
			} else {
				fmt.Printf("Syntax error: %v \n", parseError)
				//fmt.Printf("Syntax error: %v at %v\n", e, l.Position) //TODO
			}
		}
		if res, err := vm.Run(ast, env); err != nil {
			fmt.Printf("Eval error:%v\n", err)
		} else {
			debug.Printf("ENV=%#v\n", env)
			fmt.Printf("%#v\n", res)
		}
		source = ""
	}
}
