package vm

import (
	"github.com/tesujiro/exp_yacc/ast"
)

func Run(stmts []ast.Stmt, env *Env) (interface{}, error) {
	var result interface{}
	var err error
	for _, stmt := range stmts {
		result, err = runSingleStmt(stmt, env)
		if err != nil {
			return nil, err
		}
	}
	return result, err
}

func runSingleStmt(stmt ast.Stmt, env *Env) (interface{}, error) {
	switch stmt.(type) {
	case ast.AssStmt:
		return evalAssExpr(stmt.(ast.AssStmt).Left, stmt.(ast.AssStmt).Right, env)
	case ast.ExprStmt:
		return evalExpr(stmt.(ast.ExprStmt).Expr, env)
	case *ast.IfStmt:
		result, err := evalExpr(stmt.(*ast.IfStmt).If, env)
		if err != nil {
			return nil, err
		}
		if result.(bool) {
			child := env.NewEnv()
			defer child.Destroy()
			result, err = Run(stmt.(*ast.IfStmt).Then, child)
			if err != nil {
				return nil, err
			}
			//fmt.Printf("child=%#v\n", child)
			//fmt.Printf("parent=%#v\n", env)
			return result, nil
		}
		done := false
		if len(stmt.(*ast.IfStmt).ElseIf)>0 {
			for _,stmt := range stmt.(*ast.IfStmt).ElseIf {
				result, err := evalExpr(stmt.(*ast.IfStmt).If, env)
				if err != nil {
					return nil, err
				}
				if result.(bool){
					done = true
					child := env.NewEnv()
					defer child.Destroy()
					result, err = Run(stmt.(*ast.IfStmt).Then, child)
					if err != nil {
						return nil, err
					}
					return result, nil
				}
			}
		}

		if !done && len(stmt.(*ast.IfStmt).Else) > 0 {
			child := env.NewEnv()
			defer child.Destroy()
			result, err = Run(stmt.(*ast.IfStmt).Else, child)
			if err != nil {
				return nil, err
			}
		}
		return result, nil
	}
	return nil, nil
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
	default:
		// TODO:?
		//return nil,fmt.Er
	}
	return rv, nil
}
