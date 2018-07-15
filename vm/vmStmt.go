package vm

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/tesujiro/exp_yacc/ast"
)

func Run(stmts []ast.Stmt, env *Env) (interface{}, error) {
	var result interface{}
	var err error
	for _, stmt := range stmts {
		switch stmt.(type) {
		case *ast.ReturnStmt:
			result, err = runSingleStmt(stmt, env)
			if err != nil {
				return nil, err
			}
			return result, err
		default:
			result, err = runSingleStmt(stmt, env)
			if err != nil {
				return nil, err
			}
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

		// evaluate assExpr
		switch {
		case len(left) == 1 && len(right) == 1:
			return evalAssExpr(left[0], right_values[0], env)
		case len(left) > 1 && len(right) == 1:
			val := right_values[0]
			if reflect.ValueOf(val).Kind() == reflect.Interface {
				val = reflect.ValueOf(val).Elem().Interface()
			}
			if reflect.ValueOf(val).Kind() != reflect.Slice {
				return nil, errors.New("single value assign to multi values")
			} else {
				elements := reflect.ValueOf(val)
				right_values = make([]interface{}, elements.Len())
				for i := 0; i < elements.Len(); i++ {
					right_values[i] = elements.Index(i).Interface()
				}
			}
			fallthrough
		default:
			for i, expr := range left {
				if i >= len(right_values) {
					return right_values[len(right_values)-1], nil
				}
				if _, err := evalAssExpr(expr, right_values[i], env); err != nil {
					return nil, err
				}
			}
			return right_values[len(left)-1], nil
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

func evalAssExpr(lexp ast.Expr, val interface{}, env *Env) (interface{}, error) {
	switch lexp.(type) {
	case *ast.IdentExpr:
		id := lexp.(*ast.IdentExpr).Literal
		if err := env.Set(id, val); err != nil {
			env.Define(id, val)
		}
	case *ast.ItemExpr:
		index, err := evalExpr(lexp.(*ast.ItemExpr).Index, env)
		if err != nil {
			fmt.Println("ItemExpr index error") //TODO
			return nil, err
		}
		value, err := evalExpr(lexp.(*ast.ItemExpr).Value, env)
		if err != nil {
			fmt.Println("ItemExpr value error") //TODO
			return nil, err
		}

		switch reflect.TypeOf(value).Kind() {
		case reflect.Slice | reflect.Array:
			if i, ok := index.(int); !ok {
				return nil, errors.New("index cannot convert to int")
			} else {
				if i < 0 || reflect.ValueOf(value).Len() < i {
					return nil, errors.New("index out of range")
				}
				if i == reflect.ValueOf(value).Len() {
					// append val to array
					ar := reflect.Append(reflect.ValueOf(value), reflect.ValueOf(val)).Interface()
					return evalAssExpr(lexp.(*ast.ItemExpr).Value, ar, env)
				}

				// Set Val To Array
				reflect.ValueOf(value).Index(i).Set(reflect.ValueOf(val))
				return val, nil
			}
		default:
			return nil, errors.New("type " + reflect.TypeOf(value).Kind().String() + " does not support index operation")
		}

	default:
		// TODO:?
		return nil, errors.New("Invalid Operation")
	}
	return val, nil
}
