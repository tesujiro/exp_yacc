package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tesujiro/exp_yacc/parser"
	"github.com/tesujiro/exp_yacc/vm"
)

var tty *os.File

func main() {
	tty, err := os.Open(`/dev/tty`)
	if err != nil {
		panic(err)
	}
	defer tty.Close()
	//fmt.Fprintf(tty, "tty opened\n")

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
		if source == "exit" {
			break
		}

		l := new(parser.Lexer)
		l.Init(strings.NewReader(source))
		parser.Parse(l)
		//fmt.Printf("%#v\n", l.Result)
		//TODO: Error Check
		if res, err := vm.Run(l.Result, env); err != nil {
			fmt.Printf("Eval error:%v\n", err)
		} else {
			//fmt.Printf("ENV=%#v\n", env)
			fmt.Printf("%#v\n", res)
		}
		source = ""
	}
}
