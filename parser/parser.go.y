%{
package parser

import "github.com/tesujiro/exp_yacc/ast"

%}

%union{
    token ast.Token
    expr  ast.Expr
    stmt_if  ast.Stmt
    stmt  ast.Stmt
    stmts []ast.Stmt
}

%type<stmts>  program
%type<stmts>  stmts
%type<stmt>   stmt
%type<stmt_if>   stmt_if
%type<expr>   expr
%token<token> IDENT NUMBER EQEQ NEQ GE LE IF ELSE

%right '='
%left IDENT
%left EQEQ NEQ
%left '>' '<' GE LE

%left '+' '-'
%left '*' '/' '%'
%right UNARY

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
    : stmt_if
    {
        $$ = $1
    }
    | expr '=' expr
    {
        $$ = ast.AssStmt{Left: $1, Right: $3}
    }
    | expr
    {
        $$ = ast.ExprStmt{Expr: $1}
    }

stmt_if
    : stmt_if ELSE IF expr '{' stmts '}'
    {
            $$.(*ast.IfStmt).ElseIf = append($$.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: $4, Then: $6} )
    }
    | stmt_if ELSE '{' stmts '}'
    {
        if $$.(*ast.IfStmt).Else != nil {
            yylex.Error("multiple else statement")
        } else {
            $$.(*ast.IfStmt).Else = append($$.(*ast.IfStmt).Else, $4...)
        }
    }
    | IF expr '{' stmts '}'
    {
        $$ = &ast.IfStmt{If: $2, Then: $4, Else: nil}
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
    | '+' expr %prec UNARY
    {
        $$ = ast.UnaryExpr{Operator: "+", Expr:$2}
    }
    | '-' expr %prec UNARY
    {
        $$ = ast.UnaryExpr{Operator: "-", Expr:$2}
    }
    | '(' expr ')'
    {
        $$ = ast.ParentExpr{SubExpr: $2}
    }
    | expr '+' expr
    {
        $$ = ast.BinOpExpr{Left: $1, Operator: "+", Right: $3}
    }
    | expr '-' expr
    {
        $$ = ast.BinOpExpr{Left: $1, Operator: "-", Right: $3}
    }
    | expr '*' expr
    {
        $$ = ast.BinOpExpr{Left: $1, Operator: "*", Right: $3}
    }
    | expr '/' expr
    {
        $$ = ast.BinOpExpr{Left: $1, Operator: "/", Right: $3}
    }
    | expr '%' expr
    {
        $$ = ast.BinOpExpr{Left: $1, Operator: "%", Right: $3}
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
