package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/tesujiro/exp_yacc/parser"
	"github.com/tesujiro/exp_yacc/vm"
)

var debug = flag.Bool("d", false, "debug option")

func main() {
	flag.Parse()
	run()
}

func run() {
	//l := new(parser.Lexer)
	//l.Init(strings.NewReader(os.Args[1]))
	env := vm.NewEnv()
	line_scanner := bufio.NewScanner(os.Stdin) // This Scanner
	var source string

	for line_scanner.Scan() {
		source += line_scanner.Text()
		if source == "" {
			continue
		}
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
