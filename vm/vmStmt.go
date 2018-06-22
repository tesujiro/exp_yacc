package vm

import (
	"github.com/tesujiro/exp_yacc/ast"
)

func Run(stmt ast.Stmt, env *Env) (interface{}, error) {
	switch stmt.(type) {
	case ast.AssStmt:
		return evalAssExpr(stmt.(ast.AssStmt).Left, stmt.(ast.AssStmt).Right, env)
	}
	return 0, nil
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
