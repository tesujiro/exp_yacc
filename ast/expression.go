package ast

type Expr interface{}

type IdentExpr struct {
	Literal string
}

type NumExpr struct {
	Literal string
}

type BinOpExpr struct {
	Left     Expr
	Operator rune
	Right    Expr
}
