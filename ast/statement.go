package ast

type Stmt interface{}

type AssStmt struct {
	Left  Expr
	Right Expr
}

type ExprStmt struct {
	Expr Expr
}

type IfStmt struct {
	If Expr
	Then []Stmt
	Else []Stmt
	ElseIf []Stmt
}
