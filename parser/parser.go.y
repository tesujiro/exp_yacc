%{
package parser

import "github.com/tesujiro/exp_yacc/ast"

%}

%union{
    token ast.Token
    stmt_if  ast.Stmt
    stmt  ast.Stmt
    stmts []ast.Stmt
    expr  ast.Expr
    exprs  []ast.Expr
    ident_args []string
}

%type<stmts>  program
%type<stmts>  stmts
%type<stmt>   stmt
%type<stmt_if>   stmt_if
%type<expr>   expr
%type<exprs>   exprs
%type<ident_args>   ident_args
%token<token> IDENT NUMBER STRING TRUE FALSE NIL FUNC RETURN EQEQ NEQ GE LE IF ELSE

%right '='
%left IDENT
%left EQEQ NEQ
%left '>' '<' GE LE

%left '+' '-'
%left '*' '/' '%'
%right UNARY

%%

program
    : opt_term
    {
        $$ = nil
    }
    | stmts opt_term
    {
        $$ = $1
        yylex.(*Lexer).result = $$
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
    : IF expr '{' program '}'
    {
        $$ = &ast.IfStmt{If: $2, Then: $4, Else: nil}
    }
    | stmt_if ELSE IF expr '{' program '}'
    {
            $$.(*ast.IfStmt).ElseIf = append($$.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: $4, Then: $6} )
    }
    | stmt_if ELSE '{' program '}'
    {
        if $$.(*ast.IfStmt).Else != nil {
            yylex.Error("multiple else statement")
        } else {
            //$$.(*ast.IfStmt).Else = append($$.(*ast.IfStmt).Else, $4...)
            $$.(*ast.IfStmt).Else = $4
        }
    }
    | RETURN exprs
    {
        $$ = &ast.ReturnStmt{Exprs: $2}
    }

exprs
    : /* Nothing */
    {
        $$ = nil
    }
    | expr
    {
        $$ = []ast.Expr{$1}
    }
    | exprs ',' opt_newLines expr
    {
        $$ = append($1,$4)
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
    | STRING
    {
        $$ = ast.StringExpr{Literal: $1.Literal}
    }
    | TRUE
    {
        $$ = &ast.ConstExpr{Literal: $1.Literal}
    }
    | FALSE
    {
        $$ = &ast.ConstExpr{Literal: $1.Literal}
    }
    | NIL
    {
        $$ = &ast.ConstExpr{Literal: $1.Literal}
    }
    | IDENT '(' exprs ')'
    {
        $$ = &ast.CallExpr{Name: $1.Literal, SubExprs:$3}
    }
    | FUNC IDENT '(' ident_args ')''{' program '}'
    {
    	$$ = &ast.FuncExpr{Name: $2.Literal, Args: $4, Stmts: $7}
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

ident_args
    : /* nothing */
    {
    	$$ = []string{}
    }
    | IDENT
    {
    	$$ = []string{$1.Literal}
    }
    | ident_args ',' opt_newLines IDENT
    {
    	$$ = append($1,$4.Literal)
    }

opt_term
    : /* nothing */
    | term

term
    : ';' newLines
    | newLines
    | ';'

opt_newLines
    : /* nothing */
    | newLines

newLines
    : newLine
    | newLines newLine

newLine
    : '\n'

%%
