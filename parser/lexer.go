package parser

import (
	"errors"
	"text/scanner"

	"github.com/tesujiro/exp_yacc/ast"
)

type Lexer struct {
	scanner.Scanner
	result []ast.Stmt
	err    error
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
	l.err = errors.New(e)
	//l.position = l.Position //TODO
}

func Parse(yylex yyLexer) ([]ast.Stmt, error) {
	l := yylex.(*Lexer)
	if yyParse(yylex) != 0 {
		return nil, l.err
	}
	return l.result, nil
}
