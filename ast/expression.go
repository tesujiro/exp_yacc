package ast

type Expr interface{}

type IdentExpr struct {
	Literal string
}

type NumExpr struct {
	Literal string
}

type StringExpr struct {
	Literal string
}

type ParentExpr struct {
	SubExpr Expr
}

type UnaryExpr struct {
	Operator string
	Expr     Expr
}

type BinOpExpr struct {
	Left     Expr
	Operator string
	Right    Expr
}

type CallExpr struct {
	Name     string
	SubExprs []Expr
}
