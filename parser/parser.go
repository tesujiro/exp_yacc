// Code generated by goyacc - DO NOT EDIT.

package parser

import __yyfmt__ "fmt"

import "github.com/tesujiro/exp_yacc/ast"

type yySymType struct {
	yys        int
	token      ast.Token
	stmt_if    ast.Stmt
	stmt       ast.Stmt
	stmts      []ast.Stmt
	expr       ast.Expr
	exprs      []ast.Expr
	ident_args []string
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault = 57363
	yyEofCode = 57344
	ANDAND    = 57360
	ELSE      = 57359
	EQEQ      = 57354
	FALSE     = 57350
	FUNC      = 57352
	GE        = 57356
	IDENT     = 57346
	IF        = 57358
	LE        = 57357
	NEQ       = 57355
	NIL       = 57351
	NUMBER    = 57347
	OROR      = 57361
	RETURN    = 57353
	STRING    = 57348
	TRUE      = 57349
	UNARY     = 57362
	yyErrCode = 57345

	yyMaxDepth = 200
	yyTabOfs   = -58
)

var (
	yyPrec = map[int]int{
		'=':    0,
		OROR:   1,
		ANDAND: 2,
		IDENT:  3,
		EQEQ:   4,
		NEQ:    4,
		'>':    5,
		'<':    5,
		GE:     5,
		LE:     5,
		'+':    6,
		'-':    6,
		'*':    7,
		'/':    7,
		'%':    7,
		UNARY:  8,
	}

	yyXLAT = map[int]int{
		40:    0,  // '(' (84x)
		43:    1,  // '+' (82x)
		45:    2,  // '-' (82x)
		123:   3,  // '{' (79x)
		125:   4,  // '}' (69x)
		44:    5,  // ',' (68x)
		10:    6,  // '\n' (62x)
		57344: 7,  // $end (56x)
		59:    8,  // ';' (50x)
		57346: 9,  // IDENT (50x)
		61:    10, // '=' (47x)
		57350: 11, // FALSE (45x)
		57352: 12, // FUNC (45x)
		57351: 13, // NIL (45x)
		57347: 14, // NUMBER (45x)
		57348: 15, // STRING (45x)
		57349: 16, // TRUE (45x)
		41:    17, // ')' (42x)
		37:    18, // '%' (37x)
		42:    19, // '*' (37x)
		47:    20, // '/' (37x)
		60:    21, // '<' (37x)
		62:    22, // '>' (37x)
		91:    23, // '[' (37x)
		57360: 24, // ANDAND (37x)
		57359: 25, // ELSE (37x)
		57354: 26, // EQEQ (37x)
		57356: 27, // GE (37x)
		57357: 28, // LE (37x)
		57355: 29, // NEQ (37x)
		57361: 30, // OROR (37x)
		93:    31, // ']' (34x)
		57364: 32, // expr (29x)
		57358: 33, // IF (16x)
		57367: 34, // newLine (15x)
		57353: 35, // RETURN (15x)
		57368: 36, // newLines (12x)
		57365: 37, // exprs (8x)
		57370: 38, // opt_term (7x)
		57375: 39, // term (7x)
		57371: 40, // program (6x)
		57374: 41, // stmts (6x)
		57369: 42, // opt_newLines (4x)
		57366: 43, // ident_args (2x)
		57372: 44, // stmt (2x)
		57373: 45, // stmt_if (2x)
		57363: 46, // $default (0x)
		57345: 47, // error (0x)
		57362: 48, // UNARY (0x)
	}

	yySymNames = []string{
		"'('",
		"'+'",
		"'-'",
		"'{'",
		"'}'",
		"','",
		"'\\n'",
		"$end",
		"';'",
		"IDENT",
		"'='",
		"FALSE",
		"FUNC",
		"NIL",
		"NUMBER",
		"STRING",
		"TRUE",
		"')'",
		"'%'",
		"'*'",
		"'/'",
		"'<'",
		"'>'",
		"'['",
		"ANDAND",
		"ELSE",
		"EQEQ",
		"GE",
		"LE",
		"NEQ",
		"OROR",
		"']'",
		"expr",
		"IF",
		"newLine",
		"RETURN",
		"newLines",
		"exprs",
		"opt_term",
		"term",
		"program",
		"stmts",
		"opt_newLines",
		"ident_args",
		"stmt",
		"stmt_if",
		"$default",
		"error",
		"UNARY",
	}

	yyTokenLiteralStrings = map[int]string{}

	yyReductions = map[int]struct{ xsym, components int }{
		0:  {0, 1},
		1:  {40, 1},
		2:  {40, 2},
		3:  {41, 2},
		4:  {41, 3},
		5:  {44, 3},
		6:  {44, 3},
		7:  {44, 1},
		8:  {44, 1},
		9:  {45, 5},
		10: {45, 7},
		11: {45, 5},
		12: {45, 2},
		13: {37, 0},
		14: {37, 1},
		15: {37, 4},
		16: {32, 1},
		17: {32, 1},
		18: {32, 1},
		19: {32, 1},
		20: {32, 1},
		21: {32, 1},
		22: {32, 4},
		23: {32, 8},
		24: {32, 4},
		25: {32, 7},
		26: {32, 5},
		27: {32, 4},
		28: {32, 4},
		29: {32, 2},
		30: {32, 2},
		31: {32, 3},
		32: {32, 3},
		33: {32, 3},
		34: {32, 3},
		35: {32, 3},
		36: {32, 3},
		37: {32, 3},
		38: {32, 3},
		39: {32, 3},
		40: {32, 3},
		41: {32, 3},
		42: {32, 3},
		43: {32, 3},
		44: {32, 3},
		45: {43, 0},
		46: {43, 1},
		47: {43, 4},
		48: {38, 0},
		49: {38, 1},
		50: {39, 2},
		51: {39, 1},
		52: {39, 1},
		53: {42, 0},
		54: {42, 1},
		55: {36, 1},
		56: {36, 2},
		57: {34, 1},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [117][]uint16{
		// 0
		{10, 10, 10, 10, 5: 10, 66, 10, 63, 10, 10, 10, 10, 10, 10, 10, 10, 33: 10, 65, 10, 64, 38: 60, 62, 59, 61},
		{7: 58},
		{87, 85, 86, 84, 57, 45, 7: 57, 9: 77, 45, 81, 83, 82, 78, 79, 80, 32: 72, 75, 35: 76, 37: 73, 44: 174, 74},
		{4: 10, 6: 66, 10, 63, 34: 65, 36: 64, 38: 69, 70},
		{9, 9, 9, 9, 9, 9, 7: 9, 9: 9, 9, 9, 9, 9, 9, 9, 9, 33: 9, 35: 9},
		// 5
		{6, 6, 6, 6, 6, 6, 66, 6, 9: 6, 6, 6, 6, 6, 6, 6, 6, 33: 6, 65, 6, 68},
		{7, 7, 7, 7, 7, 7, 66, 7, 9: 7, 7, 7, 7, 7, 7, 7, 7, 33: 7, 67, 7},
		{3, 3, 3, 3, 3, 3, 3, 3, 9: 3, 3, 3, 3, 3, 3, 3, 3, 33: 3, 35: 3},
		{1, 1, 1, 1, 1, 1, 1, 1, 9: 1, 1, 1, 1, 1, 1, 1, 1, 33: 1, 35: 1},
		{2, 2, 2, 2, 2, 2, 2, 2, 9: 2, 2, 2, 2, 2, 2, 2, 2, 33: 2, 35: 2},
		// 10
		{8, 8, 8, 8, 8, 8, 66, 8, 9: 8, 8, 8, 8, 8, 8, 8, 8, 33: 8, 67, 8},
		{4: 56, 7: 56},
		{87, 85, 86, 84, 9, 45, 7: 9, 9: 77, 45, 81, 83, 82, 78, 79, 80, 32: 72, 75, 35: 76, 37: 73, 44: 71, 74},
		{4: 54, 6: 54, 54, 54},
		{89, 94, 95, 4: 50, 44, 50, 50, 50, 10: 172, 18: 98, 96, 97, 102, 101, 90, 93, 26: 99, 103, 104, 100, 92},
		// 15
		{5: 120, 10: 170},
		{4: 51, 6: 51, 51, 51, 25: 161},
		{87, 85, 86, 84, 9: 77, 11: 81, 83, 82, 78, 79, 80, 32: 157},
		{87, 85, 86, 84, 45, 45, 45, 45, 45, 77, 11: 81, 83, 82, 78, 79, 80, 25: 45, 32: 118, 37: 156},
		{150, 42, 42, 42, 42, 42, 42, 42, 42, 10: 42, 17: 42, 42, 42, 42, 42, 42, 151, 42, 42, 42, 42, 42, 42, 42, 42},
		// 20
		{41, 41, 41, 41, 41, 41, 41, 41, 41, 10: 41, 17: 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41},
		{40, 40, 40, 40, 40, 40, 40, 40, 40, 10: 40, 17: 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40},
		{39, 39, 39, 39, 39, 39, 39, 39, 39, 10: 39, 17: 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39},
		{38, 38, 38, 38, 38, 38, 38, 38, 38, 10: 38, 17: 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38},
		{37, 37, 37, 37, 37, 37, 37, 37, 37, 10: 37, 17: 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37},
		// 25
		{134, 9: 133},
		{5, 5, 5, 5, 5, 5, 66, 9: 5, 11: 5, 5, 5, 5, 5, 5, 34: 65, 36: 123, 42: 129},
		{87, 85, 86, 84, 9: 77, 11: 81, 83, 82, 78, 79, 80, 32: 128},
		{87, 85, 86, 84, 9: 77, 11: 81, 83, 82, 78, 79, 80, 32: 127},
		{87, 85, 86, 84, 9: 77, 11: 81, 83, 82, 78, 79, 80, 32: 88},
		// 30
		{89, 94, 95, 17: 91, 98, 96, 97, 102, 101, 90, 93, 26: 99, 103, 104, 100, 92},
		{87, 85, 86, 84, 5: 45, 9: 77, 11: 81, 83, 82, 78, 79, 80, 45, 32: 118, 37: 125},
		{87, 85, 86, 84, 5: 45, 9: 77, 11: 81, 83, 82, 78, 79, 80, 31: 45, 118, 37: 119},
		{27, 27, 27, 27, 27, 27, 27, 27, 27, 10: 27, 17: 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27},
		{87, 85, 86, 84, 9: 77, 11: 81, 83, 82, 78, 79, 80, 32: 117},
		// 35
		{87, 85, 86, 84, 9: 77, 11: 81, 83, 82, 78, 79, 80, 32: 116},
		{87, 85, 86, 84, 9: 77, 11: 81, 83, 82, 78, 79, 80, 32: 115},
		{87, 85, 86, 84, 9: 77, 11: 81, 83, 82, 78, 79, 80, 32: 114},
		{87, 85, 86, 84, 9: 77, 11: 81, 83, 82, 78, 79, 80, 32: 113},
		{87, 85, 86, 84, 9: 77, 11: 81, 83, 82, 78, 79, 80, 32: 112},
		// 40
		{87, 85, 86, 84, 9: 77, 11: 81, 83, 82, 78, 79, 80, 32: 111},
		{87, 85, 86, 84, 9: 77, 11: 81, 83, 82, 78, 79, 80, 32: 110},
		{87, 85, 86, 84, 9: 77, 11: 81, 83, 82, 78, 79, 80, 32: 109},
		{87, 85, 86, 84, 9: 77, 11: 81, 83, 82, 78, 79, 80, 32: 108},
		{87, 85, 86, 84, 9: 77, 11: 81, 83, 82, 78, 79, 80, 32: 107},
		// 45
		{87, 85, 86, 84, 9: 77, 11: 81, 83, 82, 78, 79, 80, 32: 106},
		{87, 85, 86, 84, 9: 77, 11: 81, 83, 82, 78, 79, 80, 32: 105},
		{89, 94, 95, 14, 14, 14, 14, 14, 14, 10: 14, 17: 14, 98, 96, 97, 14, 14, 90, 14, 14, 14, 14, 14, 14, 14, 14},
		{89, 94, 95, 15, 15, 15, 15, 15, 15, 10: 15, 17: 15, 98, 96, 97, 15, 15, 90, 15, 15, 15, 15, 15, 15, 15, 15},
		{89, 94, 95, 16, 16, 16, 16, 16, 16, 10: 16, 17: 16, 98, 96, 97, 16, 16, 90, 16, 16, 16, 16, 16, 16, 16, 16},
		// 50
		{89, 94, 95, 17, 17, 17, 17, 17, 17, 10: 17, 17: 17, 98, 96, 97, 17, 17, 90, 17, 17, 17, 17, 17, 17, 17, 17},
		{89, 94, 95, 18, 18, 18, 18, 18, 18, 10: 18, 17: 18, 98, 96, 97, 102, 101, 90, 18, 18, 18, 103, 104, 18, 18, 18},
		{89, 94, 95, 19, 19, 19, 19, 19, 19, 10: 19, 17: 19, 98, 96, 97, 102, 101, 90, 19, 19, 19, 103, 104, 19, 19, 19},
		{89, 20, 20, 20, 20, 20, 20, 20, 20, 10: 20, 17: 20, 20, 20, 20, 20, 20, 90, 20, 20, 20, 20, 20, 20, 20, 20},
		{89, 21, 21, 21, 21, 21, 21, 21, 21, 10: 21, 17: 21, 21, 21, 21, 21, 21, 90, 21, 21, 21, 21, 21, 21, 21, 21},
		// 55
		{89, 22, 22, 22, 22, 22, 22, 22, 22, 10: 22, 17: 22, 22, 22, 22, 22, 22, 90, 22, 22, 22, 22, 22, 22, 22, 22},
		{89, 23, 23, 23, 23, 23, 23, 23, 23, 10: 23, 17: 23, 98, 96, 97, 23, 23, 90, 23, 23, 23, 23, 23, 23, 23, 23},
		{89, 24, 24, 24, 24, 24, 24, 24, 24, 10: 24, 17: 24, 98, 96, 97, 24, 24, 90, 24, 24, 24, 24, 24, 24, 24, 24},
		{89, 94, 95, 25, 25, 25, 25, 25, 25, 10: 25, 17: 25, 98, 96, 97, 102, 101, 90, 25, 25, 99, 103, 104, 100, 25, 25},
		{89, 94, 95, 26, 26, 26, 26, 26, 26, 10: 26, 17: 26, 98, 96, 97, 102, 101, 90, 93, 26, 99, 103, 104, 100, 26, 26},
		// 60
		{89, 94, 95, 4: 44, 44, 44, 44, 44, 17: 44, 98, 96, 97, 102, 101, 90, 93, 44, 99, 103, 104, 100, 92, 44},
		{5: 120, 31: 121},
		{5, 5, 5, 5, 6: 66, 9: 5, 11: 5, 5, 5, 5, 5, 5, 34: 65, 36: 123, 42: 122},
		{30, 30, 30, 30, 30, 30, 30, 30, 30, 10: 30, 17: 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30},
		{87, 85, 86, 84, 9: 77, 11: 81, 83, 82, 78, 79, 80, 32: 124},
		// 65
		{4, 4, 4, 4, 4, 4, 66, 9: 4, 11: 4, 4, 4, 4, 4, 4, 34: 67},
		{89, 94, 95, 4: 43, 43, 43, 43, 43, 10: 43, 17: 43, 98, 96, 97, 102, 101, 90, 93, 43, 99, 103, 104, 100, 92, 43},
		{5: 120, 17: 126},
		{34, 34, 34, 34, 34, 34, 34, 34, 34, 10: 34, 17: 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34},
		{89, 28, 28, 28, 28, 28, 28, 28, 28, 10: 28, 17: 28, 28, 28, 28, 28, 28, 90, 28, 28, 28, 28, 28, 28, 28, 28},
		// 70
		{89, 29, 29, 29, 29, 29, 29, 29, 29, 10: 29, 17: 29, 29, 29, 29, 29, 29, 90, 29, 29, 29, 29, 29, 29, 29, 29},
		{87, 85, 86, 84, 45, 45, 45, 9: 77, 11: 81, 83, 82, 78, 79, 80, 32: 118, 37: 130},
		{4: 5, 120, 66, 34: 65, 36: 123, 42: 131},
		{4: 132},
		{32, 32, 32, 32, 32, 32, 32, 32, 32, 10: 32, 17: 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32},
		// 75
		{144},
		{5: 13, 9: 136, 17: 13, 43: 135},
		{5: 138, 17: 137},
		{5: 12, 17: 12},
		{3: 141},
		// 80
		{6: 66, 9: 5, 34: 65, 36: 123, 42: 139},
		{9: 140},
		{5: 11, 17: 11},
		{10, 10, 10, 10, 10, 10, 66, 8: 63, 10, 10, 10, 10, 10, 10, 10, 10, 33: 10, 65, 10, 64, 38: 60, 62, 142, 61},
		{4: 143},
		// 85
		{33, 33, 33, 33, 33, 33, 33, 33, 33, 10: 33, 17: 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33},
		{5: 13, 9: 136, 17: 13, 43: 145},
		{5: 138, 17: 146},
		{3: 147},
		{10, 10, 10, 10, 10, 10, 66, 8: 63, 10, 10, 10, 10, 10, 10, 10, 10, 33: 10, 65, 10, 64, 38: 60, 62, 148, 61},
		// 90
		{4: 149},
		{35, 35, 35, 35, 35, 35, 35, 35, 35, 10: 35, 17: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35},
		{87, 85, 86, 84, 5: 45, 9: 77, 11: 81, 83, 82, 78, 79, 80, 45, 32: 118, 37: 154},
		{87, 85, 86, 84, 9: 77, 11: 81, 83, 82, 78, 79, 80, 32: 152},
		{89, 94, 95, 18: 98, 96, 97, 102, 101, 90, 93, 26: 99, 103, 104, 100, 92, 153},
		// 95
		{31, 31, 31, 31, 31, 31, 31, 31, 31, 10: 31, 17: 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31},
		{5: 120, 17: 155},
		{36, 36, 36, 36, 36, 36, 36, 36, 36, 10: 36, 17: 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36},
		{4: 46, 120, 46, 46, 46, 25: 46},
		{89, 94, 95, 158, 18: 98, 96, 97, 102, 101, 90, 93, 26: 99, 103, 104, 100, 92},
		// 100
		{10, 10, 10, 10, 10, 10, 66, 8: 63, 10, 10, 10, 10, 10, 10, 10, 10, 33: 10, 65, 10, 64, 38: 60, 62, 159, 61},
		{4: 160},
		{4: 49, 6: 49, 49, 49, 25: 49},
		{3: 163, 33: 162},
		{87, 85, 86, 84, 9: 77, 11: 81, 83, 82, 78, 79, 80, 32: 166},
		// 105
		{10, 10, 10, 10, 10, 10, 66, 8: 63, 10, 10, 10, 10, 10, 10, 10, 10, 33: 10, 65, 10, 64, 38: 60, 62, 164, 61},
		{4: 165},
		{4: 47, 6: 47, 47, 47, 25: 47},
		{89, 94, 95, 167, 18: 98, 96, 97, 102, 101, 90, 93, 26: 99, 103, 104, 100, 92},
		{10, 10, 10, 10, 10, 10, 66, 8: 63, 10, 10, 10, 10, 10, 10, 10, 10, 33: 10, 65, 10, 64, 38: 60, 62, 168, 61},
		// 110
		{4: 169},
		{4: 48, 6: 48, 48, 48, 25: 48},
		{87, 85, 86, 84, 45, 45, 45, 45, 45, 77, 11: 81, 83, 82, 78, 79, 80, 32: 118, 37: 171},
		{4: 52, 120, 52, 52, 52},
		{87, 85, 86, 84, 9: 77, 11: 81, 83, 82, 78, 79, 80, 32: 173},
		// 115
		{89, 94, 95, 4: 53, 6: 53, 53, 53, 18: 98, 96, 97, 102, 101, 90, 93, 26: 99, 103, 104, 100, 92},
		{4: 55, 6: 55, 55, 55},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	if c < 0x7f {
		return __yyfmt__.Sprintf("%q", c)
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), lval: %+v\n", yySymName(n), n, n, lval)
	}
	return n
}

func yyParse(yylex yyLexer) int {
	const yyError = 47

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yylval.yys = yystate
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := yyTokenLiteralStrings[yychar]
				if ls == "" {
					ls = yySymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 1:
		{
			yyVAL.stmts = nil
		}
	case 2:
		{
			yyVAL.stmts = yyS[yypt-1].stmts
			yylex.(*Lexer).result = yyVAL.stmts
		}
	case 3:
		{
			yyVAL.stmts = []ast.Stmt{yyS[yypt-0].stmt}
		}
	case 4:
		{
			yyVAL.stmts = append(yyS[yypt-2].stmts, yyS[yypt-0].stmt)
		}
	case 5:
		{
			yyVAL.stmt = &ast.AssStmt{Left: []ast.Expr{yyS[yypt-2].expr}, Right: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 6:
		{
			yyVAL.stmt = &ast.AssStmt{Left: yyS[yypt-2].exprs, Right: yyS[yypt-0].exprs}
		}
	case 7:
		{
			yyVAL.stmt = yyS[yypt-0].stmt_if
		}
	case 8:
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr}
		}
	case 9:
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-3].expr, Then: yyS[yypt-1].stmts, Else: nil}
		}
	case 10:
		{
			yyVAL.stmt_if.(*ast.IfStmt).ElseIf = append(yyVAL.stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyS[yypt-3].expr, Then: yyS[yypt-1].stmts})
		}
	case 11:
		{
			if yyVAL.stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				//$$.(*ast.IfStmt).Else = append($$.(*ast.IfStmt).Else, $4...)
				yyVAL.stmt_if.(*ast.IfStmt).Else = yyS[yypt-1].stmts
			}
		}
	case 12:
		{
			yyVAL.stmt_if = &ast.ReturnStmt{Exprs: yyS[yypt-0].exprs}
		}
	case 13:
		{
			yyVAL.exprs = nil
		}
	case 14:
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 15:
		{
			yyVAL.exprs = append(yyS[yypt-3].exprs, yyS[yypt-0].expr)
		}
	case 16:
		{
			yyVAL.expr = &ast.IdentExpr{Literal: yyS[yypt-0].token.Literal}
		}
	case 17:
		{
			yyVAL.expr = &ast.NumExpr{Literal: yyS[yypt-0].token.Literal}
		}
	case 18:
		{
			yyVAL.expr = &ast.StringExpr{Literal: yyS[yypt-0].token.Literal}
		}
	case 19:
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyS[yypt-0].token.Literal}
		}
	case 20:
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyS[yypt-0].token.Literal}
		}
	case 21:
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyS[yypt-0].token.Literal}
		}
	case 22:
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].token.Literal, SubExprs: yyS[yypt-1].exprs}
		}
	case 23:
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-6].token.Literal, Args: yyS[yypt-4].ident_args, Stmts: yyS[yypt-1].stmts}
		}
	case 24:
		{
			yyVAL.expr = &ast.AnonymousCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
		}
	case 25:
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].ident_args, Stmts: yyS[yypt-1].stmts}
		}
	case 26:
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-2].exprs}
		}
	case 27:
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Literal: yyS[yypt-3].token.Literal}, Index: yyS[yypt-1].expr}
		}
	case 28:
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].exprs}
		}
	case 29:
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "+", Expr: yyS[yypt-0].expr}
		}
	case 30:
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyS[yypt-0].expr}
		}
	case 31:
		{
			yyVAL.expr = &ast.ParentExpr{SubExpr: yyS[yypt-1].expr}
		}
	case 32:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "||", Right: yyS[yypt-0].expr}
		}
	case 33:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "&&", Right: yyS[yypt-0].expr}
		}
	case 34:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "+", Right: yyS[yypt-0].expr}
		}
	case 35:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "-", Right: yyS[yypt-0].expr}
		}
	case 36:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "*", Right: yyS[yypt-0].expr}
		}
	case 37:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "/", Right: yyS[yypt-0].expr}
		}
	case 38:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "%", Right: yyS[yypt-0].expr}
		}
	case 39:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "==", Right: yyS[yypt-0].expr}
		}
	case 40:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "!=", Right: yyS[yypt-0].expr}
		}
	case 41:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: ">", Right: yyS[yypt-0].expr}
		}
	case 42:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "<", Right: yyS[yypt-0].expr}
		}
	case 43:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: ">=", Right: yyS[yypt-0].expr}
		}
	case 44:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "<=", Right: yyS[yypt-0].expr}
		}
	case 45:
		{
			yyVAL.ident_args = []string{}
		}
	case 46:
		{
			yyVAL.ident_args = []string{yyS[yypt-0].token.Literal}
		}
	case 47:
		{
			yyVAL.ident_args = append(yyS[yypt-3].ident_args, yyS[yypt-0].token.Literal)
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
