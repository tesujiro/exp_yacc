package vm

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/tesujiro/exp_yacc/ast"
)

func toInt64(val interface{}) int64 {
	switch reflect.TypeOf(val).Kind() {
	case reflect.Float64, reflect.Float32:
		return int64(val.(float64))
	}
	i, _ := val.(int64)
	return i
}

func toFloat64(val interface{}) float64 {
	switch reflect.TypeOf(val).Kind() {
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
		return float64(val.(int64))
	}
	f, _ := val.(float64)
	return f
}

func evalExpr(expr ast.Expr, env *Env) (interface{}, error) {
	switch expr.(type) {
	case ast.IdentExpr:
		id := expr.(ast.IdentExpr).Literal
		if val, err := env.Get(id); err != nil {
			return nil, err
		} else {
			return val, nil
		}
	case ast.NumExpr:
		lit := expr.(ast.NumExpr).Literal
		if strings.Contains(lit, ".") {
			if f, err := strconv.ParseFloat(lit, 64); err != nil {
				return 0.0, err
			} else {
				return f, nil
			}
		}
		if i, err := strconv.ParseInt(lit, 10, 64); err != nil {
			return 0, err
		} else {
			return i, nil
		}
	case ast.ParentExpr:
		sub := expr.(ast.ParentExpr).SubExpr
		return evalExpr(sub, env)
	case ast.UnaryExpr:
		var val interface{}
		var err error
		if val, err = evalExpr(expr.(ast.UnaryExpr).Expr, env); err != nil {
			return nil, err
		}
		switch expr.(ast.UnaryExpr).Operator {
		case "+":
			return val, nil
		case "-":
			kind := reflect.TypeOf(val).Kind()
			switch kind {
			case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
				return -1 * toInt64(val), nil
			case reflect.Float64, reflect.Float32:
				return -1 * toFloat64(val), nil
			}
		}
	case ast.BinOpExpr:
		var left, right interface{}
		var err error
		if left, err = evalExpr(expr.(ast.BinOpExpr).Left, env); err != nil {
			return nil, err
		}
		if right, err = evalExpr(expr.(ast.BinOpExpr).Right, env); err != nil {
			return nil, err
		}
		/*
			//TODO: checking type  "a==1"
			if left == nil && right == nil {
				return 0, nil
			} else if left == nil {
				return right, nil
			} else if right == nil {
				return left, nil
			}
		*/
		switch expr.(ast.BinOpExpr).Operator {
		case "==":
			return left == right, nil
		case "!=":
			return left != right, nil
		case ">":
			//fmt.Printf("toFloat(left)=%#v\n", toFloat64(left))
			//fmt.Printf("toFloat(right)=%#v\n", toFloat64(right))
			return toFloat64(left) > toFloat64(right), nil
		case ">=":
			return toFloat64(left) >= toFloat64(right), nil
		case "<":
			return toFloat64(left) < toFloat64(right), nil
		case "<=":
			return toFloat64(left) <= toFloat64(right), nil
		case "+":
			l_kind := reflect.TypeOf(left).Kind()
			r_kind := reflect.TypeOf(right).Kind()
			switch {
			case l_kind == reflect.Int64 && r_kind == reflect.Int64:
				return toInt64(left) + toInt64(right), nil
			default:
				return toFloat64(left) + toFloat64(right), nil
			}
		case "-":
			l_kind := reflect.TypeOf(left).Kind()
			r_kind := reflect.TypeOf(right).Kind()
			switch {
			case l_kind == reflect.Int64 && r_kind == reflect.Int64:
				return toInt64(left) - toInt64(right), nil
			default:
				return toFloat64(left) - toFloat64(right), nil
			}
		case "*":
			l_kind := reflect.TypeOf(left).Kind()
			r_kind := reflect.TypeOf(right).Kind()
			switch {
			case l_kind == reflect.Int64 && r_kind == reflect.Int64:
				return toInt64(left) * toInt64(right), nil
			default:
				return toFloat64(left) * toFloat64(right), nil
			}
		case "/":
			l_kind := reflect.TypeOf(left).Kind()
			r_kind := reflect.TypeOf(right).Kind()
			if right == int64(0) {
				return nil, fmt.Errorf("devision by zero")
			}
			switch {
			case l_kind == reflect.Int64 && r_kind == reflect.Int64:
				return toInt64(left) / toInt64(right), nil
			default:
				return toFloat64(left) / toFloat64(right), nil
			}
		case "%":
			return toInt64(left) % toInt64(right), nil
		}
	}
	return 0, nil
}
