%{
package parser

import "github.com/tesujiro/exp_yacc/ast"

%}

%union{
    token ast.Token
    expr  ast.Expr
    stmt  ast.Stmt
    stmts []ast.Stmt
}

%type<stmts>  program
%type<stmts>  stmts
%type<stmt>   stmt
%type<expr>   expr
%token<token> IDENT NUMBER EQEQ NEQ GE LE

%left ';' '\n'
%right '='
%left IDENT
%left EQEQ NEQ
%left '>' '<' GE LE

%left '+' '-'

%%

program
    : stmts opt_term
    {
        $$ = $1
        yylex.(*Lexer).Result = $$
    }

stmts
    : opt_term stmt
    {
        $$ = []ast.Stmt{$2}
    }
    | stmts term stmt
    {
        $$ = append($1,$3)
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
        $$ = ast.BinOpExpr{Left: $1, Operator: "+", Right: $3}
    }
    | expr '-' expr
    {
        $$ = ast.BinOpExpr{Left: $1, Operator: "-", Right: $3}
    }
    | expr EQEQ expr
    {
        $$ = ast.BinOpExpr{Left: $1, Operator: "==", Right: $3}
    }
    | expr NEQ expr
    {
        $$ = ast.BinOpExpr{Left: $1, Operator: "!=", Right: $3}
    }
    | expr '>' expr
    {
        $$ = ast.BinOpExpr{Left: $1, Operator: ">", Right: $3}
    }
    | expr '<' expr
    {
        $$ = ast.BinOpExpr{Left: $1, Operator: "<", Right: $3}
    }
    | expr GE expr
    {
        $$ = ast.BinOpExpr{Left: $1, Operator: ">=", Right: $3}
    }
    | expr LE expr
    {
        $$ = ast.BinOpExpr{Left: $1, Operator: "<=", Right: $3}
    }

opt_term
    : /* nothing */
    | term

term
    : ';' newLines
    | newLines
    | ';'

newLines
    : newLine
    | newLines newLine

newLine
    : '\n'

%%
