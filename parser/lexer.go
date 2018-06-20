package parser

import (
	"strconv"
	"text/scanner"

	"github.com/tesujiro/exp_yacc/ast"
)

type Lexer struct {
	scanner.Scanner
	Result ast.Expression
}

func (l *Lexer) Lex(lval *yySymType) int {
	token := int(l.Scan())
	//if token == scanner.Int {
	//}
	switch token {
	case scanner.Int:
		token = NUMBER
	case scanner.Float:
		token = NUMBER
	}
	lval.token = ast.Token{Token: token, Literal: l.TokenText()}
	return token
}

func (l *Lexer) Error(e string) {
	panic(e)
}

func Eval(e ast.Expression) int {
	switch e.(type) {
	case ast.BinOpExpr:
		left := Eval(e.(ast.BinOpExpr).Left)
		right := Eval(e.(ast.BinOpExpr).Right)
		switch e.(ast.BinOpExpr).Operator {
		case '+':
			return left + right
		case '-':
			return left - right
		}
	case ast.NumExpr:
		num, _ := strconv.Atoi(e.(ast.NumExpr).Literal)
		return num
	}
	return 0
}

func Parse(yylex yyLexer) int {
	return yyParse(yylex)
}
