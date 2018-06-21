package ast

type Statement interface{}

type AssStmt struct {
	Left  Expression
	Right Expression
}

type ExprStmt struct {
	Expr Expression
}
