package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/tesujiro/exp_yacc/parser"
	"github.com/tesujiro/exp_yacc/vm"
)

func main() {
	l := new(parser.Lexer)
	l.Init(strings.NewReader(os.Args[1]))
	env := vm.NewEnv()

	//l.Init(os.Stdin)
	//for parser.Parse(l) >= 0 {
	parser.Parse(l)
	fmt.Printf("%#v\n", l.Result)
	if res, err := vm.Run(l.Result, env); err != nil {
		fmt.Printf("Eval error:%v\n", err)
	} else {
		fmt.Printf("ENV=%#v\n", env)
		fmt.Printf("RESULT=%#v\n", res)
	}
}
