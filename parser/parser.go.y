%{
package parser

%}

%union{
    token Token
    expr  Expression
}

%type<expr> program
%type<expr> expr
%token<token> NUMBER

%left '+'

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
        $$ = NumExpr{literal: $1.literal}
    }
    | expr '+' expr
    {
        $$ = BinOpExpr{left: $1, operator: '+', right: $3}
    }

%%
