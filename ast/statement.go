package ast

type Statement interface{}

type AssStmt struct {
	Left  Expr
	Right Expr
}

type ExprStmt struct {
	Expr Expr
}
