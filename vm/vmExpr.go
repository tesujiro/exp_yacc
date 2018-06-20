package vm

import (
	"strconv"

	"github.com/tesujiro/exp_yacc/ast"
)

func Eval(e ast.Expression) int {
	switch e.(type) {
	case ast.BinOpExpr:
		left := Eval(e.(ast.BinOpExpr).Left)
		right := Eval(e.(ast.BinOpExpr).Right)
		switch e.(ast.BinOpExpr).Operator {
		case '+':
			return left + right
		case '-':
			return left - right
		}
	case ast.NumExpr:
		num, _ := strconv.Atoi(e.(ast.NumExpr).Literal)
		return num
	}
	return 0
}
