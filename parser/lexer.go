package parser

import (
	"fmt"
	"text/scanner"

	"github.com/tesujiro/exp_yacc/ast"
)

type Lexer struct {
	scanner.Scanner
	//Result ast.Expr
	Result []ast.Stmt
}

// opName is a correction of operation names.
var opName = map[string]int{
	"if":   IF,
	"else": ELSE,
}

func (l *Lexer) Lex(lval *yySymType) int {
	token := int(l.Scan())
	switch token {
	case scanner.Ident:
		if name, ok := opName[l.TokenText()]; ok {
			token = name
		} else {
			token = IDENT
		}
	case scanner.Int:
		token = NUMBER
	case scanner.Float:
		token = NUMBER
	case scanner.String:
		token = STRING
		lit := l.TokenText()
		if len(lit) > 1 {
			lit = lit[1 : len(lit)-1]
		}
		lval.token = ast.Token{Token: token, Literal: lit}
		return token
	case int('='):
		switch l.Peek() {
		case '=':
			token = EQEQ
			l.Next()
		}
	case int('!'):
		switch l.Peek() {
		case '=':
			token = NEQ
			l.Next()
		}
	case int('>'):
		switch l.Peek() {
		case '=':
			token = GE
			l.Next()
		}
	case int('<'):
		switch l.Peek() {
		case '=':
			token = LE
			l.Next()
		}
	}
	lval.token = ast.Token{Token: token, Literal: l.TokenText()}
	//fmt.Printf("token=%#v\n", lval.token)
	return token
}

func (l *Lexer) Error(e string) {
	fmt.Printf("Syntax error: %v at %v\n", e, l.Position)
}

func Parse(yylex yyLexer) int {
	//fmt.Println("Parse")
	return yyParse(yylex)
}
