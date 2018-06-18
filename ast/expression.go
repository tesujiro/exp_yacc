package ast

type Expression interface{}

type NumExpr struct {
	Literal string
}

type BinOpExpr struct {
	Left     Expression
	Operator rune
	Right    Expression
}
