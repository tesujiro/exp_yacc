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
	parser.Parse(l)
	fmt.Printf("%#v\n", l.Result)
	fmt.Printf("RESULT=%#v\n", vm.Eval(l.Result))
}
