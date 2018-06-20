package vm

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/tesujiro/exp_yacc/ast"
)

func toInt64(lit interface{}) int64 {
	i, _ := lit.(int64)
	return i
}

func toFloat64(lit interface{}) float64 {
	f, _ := lit.(float64)
	return f
}

func Eval(e ast.Expression) (interface{}, error) {
	switch e.(type) {
	case ast.BinOpExpr:
		var left, right interface{}
		var err error
		if left, err = Eval(e.(ast.BinOpExpr).Left); err != nil {
			return nil, err
		}
		if right, err = Eval(e.(ast.BinOpExpr).Right); err != nil {
			return nil, err
		}
		switch e.(ast.BinOpExpr).Operator {
		case '+':
			l_kind := reflect.TypeOf(left).Kind()
			r_kind := reflect.TypeOf(right).Kind()
			switch {
			case l_kind == reflect.Int64 && r_kind == reflect.Int64:
				return toInt64(left) + toInt64(right), nil
			case l_kind == reflect.Int64 && r_kind == reflect.Float64:
				return float64(toInt64(left)) + toFloat64(right), nil
			case l_kind == reflect.Float64 && r_kind == reflect.Int64:
				return toFloat64(left) + float64(toInt64(right)), nil
			default:
				return toFloat64(left) + toFloat64(right), nil
			}
		case '-':
			l_kind := reflect.TypeOf(left).Kind()
			r_kind := reflect.TypeOf(right).Kind()
			switch {
			case l_kind == reflect.Int64 && r_kind == reflect.Int64:
				return toInt64(left) - toInt64(right), nil
			case l_kind == reflect.Int64 && r_kind == reflect.Float64:
				return float64(toInt64(left)) - toFloat64(right), nil
			case l_kind == reflect.Float64 && r_kind == reflect.Int64:
				return toFloat64(left) - float64(toInt64(right)), nil
			default:
				return toFloat64(left) - toFloat64(right), nil
			}
		}
	case ast.NumExpr:
		lit := e.(ast.NumExpr).Literal
		if strings.Contains(lit, ".") {
			if f, err := strconv.ParseFloat(lit, 64); err != nil {
				return 0.0, nil
			} else {
				return f, nil
			}
		}
		if i, err := strconv.ParseInt(lit, 10, 64); err != nil {
			return 0, nil
		} else {
			return i, nil
		}
	}
	return 0, nil
}
