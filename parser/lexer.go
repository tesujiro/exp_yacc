package parser

import (
	"errors"
	"go/scanner"
	"go/token"

	"github.com/tesujiro/exp_yacc/ast"
)

type Lexer struct {
	scanner.Scanner
	result []ast.Stmt
	err    error
}

// opName is a correction of operation names.
var opName = map[string]int{
	"true":     TRUE,
	"false":    FALSE,
	"nil":      NIL,
	"if":       IF,
	"else":     ELSE,
	"func":     FUNC,
	"return":   RETURN,
	"len":      LEN,
	"for":      FOR,
	"break":    BREAK,
	"continue": CONTINUE,
}

func (l *Lexer) Lex(lval *yySymType) (token_id int) {
	//TODO: Position
	_, tok, lit := l.Scan()
	if name, ok := opName[lit]; ok {
		token_id = name
		lval.token = ast.Token{Token: token_id, Literal: lit}
		//fmt.Printf("tok=%v\ttoken=%#v\n", tok.String(), lval.token)
		return token_id
	}
	switch tok {
	case token.IDENT:
		token_id = IDENT
	case token.INT:
		token_id = NUMBER
	case token.FLOAT:
		token_id = NUMBER
	case token.STRING:
		token_id = STRING
		if len(lit) > 1 {
			lit = lit[1 : len(lit)-1]
		}
	default:
		switch tok.String() {
		case "==":
			token_id = EQEQ
		case "!=":
			token_id = NEQ
		case ">=":
			token_id = GE
		case "<=":
			token_id = LE
		case "&&":
			token_id = ANDAND
		case "||":
			token_id = OROR
		case "++":
			token_id = PLUSPLUS
		case "--":
			token_id = MINUSMINUS
		case "+=":
			token_id = PLUSEQ
		case "-=":
			token_id = MINUSEQ
		case "*=":
			token_id = MULEQ
		case "/=":
			token_id = DIVEQ
		default:
			if len(tok.String()) == 1 {
				token_id = int(tok.String()[0])
			}
		}
	}
	lval.token = ast.Token{Token: token_id, Literal: lit}
	//fmt.Printf("tok=%v\ttoken=%#v\n", tok.String(), lval.token)
	return token_id
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
