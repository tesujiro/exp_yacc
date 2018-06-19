%{
package parser

import "github.com/tesujiro/exp_yacc/ast"

%}

%union{
    token ast.Token
    expr  ast.Expression
}

%type<expr> program
%type<expr> expr
%token<token> NUMBER

%left '+' '-'

%%

program
    : expr
    {
        $$ = $1
        yylex.(*Lexer).Result = $$
    }

expr
    : NUMBER
    {
        $$ = ast.NumExpr{Literal: $1.Literal}
    }
    | expr '+' expr
    {
        $$ = ast.BinOpExpr{Left: $1, Operator: '+', Right: $3}
    }
    | expr '-' expr
    {
        $$ = ast.BinOpExpr{Left: $1, Operator: '-', Right: $3}
    }

%%
