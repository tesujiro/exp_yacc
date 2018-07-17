package ast

type Stmt interface{}

type AssStmt struct {
	Left  []Expr
	Right []Expr
}

type ExprStmt struct {
	Expr Expr
}

type IfStmt struct {
	If     Expr
	Then   []Stmt
	Else   []Stmt
	ElseIf []Stmt
}

type ReturnStmt struct {
	Exprs []Expr
}

type LoopStmt struct {
	Stmts []Stmt
}

type BreakStmt struct {
}

type ContinueStmt struct {
}
