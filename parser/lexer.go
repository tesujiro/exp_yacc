package parser

import (
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

func Parse(yylex yyLexer) int {
	return yyParse(yylex)
}
