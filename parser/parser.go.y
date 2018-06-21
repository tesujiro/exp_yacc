%{
package parser

import "github.com/tesujiro/exp_yacc/ast"

%}

%union{
    token ast.Token
    expr  ast.Expression
    stmt  ast.Statement
}

%type<expr> program
%type<expr> expr
%type<stmt> stmt
%token<token> IDENT NUMBER

%right '='
%left IDENT
%left '+' '-'

%%

    //: expr
program
    : stmt
    {
        $$ = $1
        yylex.(*Lexer).Result = $$
    }

stmt
    : expr '=' expr
    {
        $$ = ast.AssStmt{Left: $1, Right: $3}
    }
    | expr
    {
        $$ = ast.ExprStmt{Expr: $1}
    }

expr
    : IDENT
    {
        $$ = ast.IdentExpr{Literal: $1.Literal}
    }
    | NUMBER
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
