package vm

import (
	"github.com/tesujiro/exp_yacc/ast"
)

func Run(stmts []ast.Stmt, env *Env) (interface{}, error) {
	var result interface{}
	var err error
	for _, stmt := range stmts {
		switch stmt.(type) {
		case ast.AssStmt:
			result, err = evalAssExpr(stmt.(ast.AssStmt).Left, stmt.(ast.AssStmt).Right, env)
		case ast.ExprStmt:
			result, err = evalExpr(stmt.(ast.ExprStmt).Expr, env)
		}
	}
	return result, err
}

func evalAssExpr(lexp ast.Expr, rexp ast.Expr, env *Env) (interface{}, error) {
	var rv interface{}
	var err error
	if rv, err = evalExpr(rexp, env); err != nil {
		return nil, err
	}
	if rv == nil {
		return nil, nil
	}
	switch lexp.(type) {
	case ast.IdentExpr:
		id := lexp.(ast.IdentExpr).Literal
		if err := env.Set(id, rv); err != nil {
			env.Define(id, rv)
		}
	}
	return rv, nil
}
