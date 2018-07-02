package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/tesujiro/exp_yacc/parser"
	"github.com/tesujiro/exp_yacc/vm"
)

var debug = flag.Bool("d", false, "debug option")
var file string

func main() {
	flag.Parse()
	file = flag.Arg(0)

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
	l.Init(strings.NewReader(source))
	ast, parseError := parser.Parse(l)
	if *debug {
		fmt.Printf("%#v\n", ast)
	}
	if parseError != nil {
		fmt.Printf("Syntax error: %v \n", parseError)
		//fmt.Printf("Syntax error: %v at %v\n", e, l.Position) //TODO
		return
	}
	if res, err := vm.Run(ast, env); err != nil {
		fmt.Printf("Eval error:%v\n", err)
		return
	} else {
		if *debug {
			fmt.Printf("ENV=%#v\n", env)
		}
		fmt.Printf("%#v\n", res)
	}
	return
}

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
		l.Init(strings.NewReader(source))
		ast, parseError := parser.Parse(l)
		if *debug {
			fmt.Printf("%#v\n", ast)
		}
		if parseError != nil {
			if parseError.Error() == "unexpected $end" {
				// caution: scanner.Scan() does not return "end of line" ,
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
			if *debug {
				fmt.Printf("ENV=%#v\n", env)
			}
			fmt.Printf("%#v\n", res)
		}
		source = ""
	}
}
