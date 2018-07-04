package parser

import (
	"errors"
	"text/scanner"

	"github.com/tesujiro/exp_yacc/ast"
)

type Lexer struct {
	scanner.Scanner
	result  []ast.Stmt
	err     error
	buffer  []int
	curLine int
}

// opName is a correction of operation names.
var opName = map[string]int{
	"if":   IF,
	"else": ELSE,
	"func": FUNC,
}

func (l *Lexer) getToken() int {
	if len(l.buffer) > 0 {
		token := l.buffer[0]
		l.buffer = l.buffer[1:]
		return token
	}
	token := int(l.Scan())
	if l.Position.Line > l.curLine {
		// Add '\n' to buffer
		for ; l.Position.Line > l.curLine; l.curLine++ {
			l.buffer = append(l.buffer, int('\n'))
		}
		l.buffer = append(l.buffer, int(token))
		l.curLine = l.Position.Line
		return l.getToken()
	}
	l.curLine = l.Position.Line
	return token
}

func (l *Lexer) Lex(lval *yySymType) int {
	//token := int(l.Scan())
	token := l.getToken()
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
