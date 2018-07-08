package vm

import (
	"errors"

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
	case *ast.AssStmt:
		assStmt := stmt.(*ast.AssStmt)
		left, right := assStmt.Left, assStmt.Right

		// evaluate right expressions
		right_values := make([]interface{}, len(right))
		var err error
		for i, expr := range right {
			right_values[i], err = evalExpr(expr, env)
			if err != nil {
				return nil, err
			}
		}

		//
		switch {
		case len(left) == 1 && len(right) == 1:
			return evalAssExpr(left[0], right_values[0], env)
		}
	case *ast.ExprStmt:
		return evalExpr(stmt.(*ast.ExprStmt).Expr, env)
	case *ast.IfStmt:
		child := env.NewEnv()
		defer child.Destroy()
		result, err := evalExpr(stmt.(*ast.IfStmt).If, child)
		if err != nil {
			return nil, err
		}
		if result.(bool) {
			result, err = Run(stmt.(*ast.IfStmt).Then, child)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		for _, stmt := range stmt.(*ast.IfStmt).ElseIf {
			result, err := evalExpr(stmt.(*ast.IfStmt).If, child)
			if err != nil {
				return nil, err
			}
			if result.(bool) {
				result, err = Run(stmt.(*ast.IfStmt).Then, child)
				if err != nil {
					return nil, err
				}
				return result, nil
			}
		}

		if len(stmt.(*ast.IfStmt).Else) > 0 {
			result, err = Run(stmt.(*ast.IfStmt).Else, child)
			if err != nil {
				return nil, err
			}
		}
		return result, nil
	case *ast.ReturnStmt:
		returnStmt := stmt.(*ast.ReturnStmt)
		length := len(returnStmt.Exprs)

		resultExpr := make([]interface{}, length)
		var err error
		for i, expr := range returnStmt.Exprs {
			resultExpr[i], err = evalExpr(expr, env)
			if err != nil {
				return nil, err
			}
		}

		switch length {
		case 0:
			return nil, nil
		case 1:
			return resultExpr[0], nil
		default:
			return resultExpr, nil
		}
	}
	return nil, nil
}

/*func evalAssExpr(lexp ast.Expr, rexp ast.Expr, env *Env) (interface{}, error) {
var rv interface{}
var err error
if rv, err = evalExpr(rexp, env); err != nil {
	return nil, err
}
if rv == nil {
	return nil, nil
}
*/

func evalAssExpr(lexp ast.Expr, val interface{}, env *Env) (interface{}, error) {
	switch lexp.(type) {
	case *ast.IdentExpr:
		id := lexp.(*ast.IdentExpr).Literal
		if err := env.Set(id, val); err != nil {
			env.Define(id, val)
		}
	default:
		// TODO:?
		return nil, errors.New("Invalid Operation")
	}
	return val, nil
}
