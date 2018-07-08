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
	yyDefault = 57361
	yyEofCode = 57344
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
	RETURN    = 57353
	STRING    = 57348
	TRUE      = 57349
	UNARY     = 57360
	yyErrCode = 57345

	yyMaxDepth = 200
	yyTabOfs   = -51
)

var (
	yyPrec = map[int]int{
		'=':   0,
		IDENT: 1,
		EQEQ:  2,
		NEQ:   2,
		'>':   3,
		'<':   3,
		GE:    3,
		LE:    3,
		'+':   4,
		'-':   4,
		'*':   5,
		'/':   5,
		'%':   5,
		UNARY: 6,
	}

	yyXLAT = map[int]int{
		43:    0,  // '+' (66x)
		45:    1,  // '-' (66x)
		125:   2,  // '}' (55x)
		10:    3,  // '\n' (51x)
		44:    4,  // ',' (50x)
		57344: 5,  // $end (49x)
		59:    6,  // ';' (42x)
		57346: 7,  // IDENT (41x)
		40:    8,  // '(' (39x)
		61:    9,  // '=' (39x)
		57350: 10, // FALSE (37x)
		57352: 11, // FUNC (37x)
		57351: 12, // NIL (37x)
		57347: 13, // NUMBER (37x)
		57348: 14, // STRING (37x)
		57349: 15, // TRUE (37x)
		41:    16, // ')' (31x)
		57359: 17, // ELSE (30x)
		37:    18, // '%' (29x)
		42:    19, // '*' (29x)
		47:    20, // '/' (29x)
		60:    21, // '<' (29x)
		62:    22, // '>' (29x)
		57354: 23, // EQEQ (29x)
		57356: 24, // GE (29x)
		57357: 25, // LE (29x)
		57355: 26, // NEQ (29x)
		123:   27, // '{' (26x)
		57362: 28, // expr (23x)
		57358: 29, // IF (15x)
		57353: 30, // RETURN (14x)
		57365: 31, // newLine (12x)
		57366: 32, // newLines (9x)
		57368: 33, // opt_term (6x)
		57373: 34, // term (6x)
		57363: 35, // exprs (5x)
		57369: 36, // program (5x)
		57372: 37, // stmts (5x)
		57367: 38, // opt_newLines (2x)
		57370: 39, // stmt (2x)
		57371: 40, // stmt_if (2x)
		57364: 41, // ident_args (1x)
		57361: 42, // $default (0x)
		57345: 43, // error (0x)
		57360: 44, // UNARY (0x)
	}

	yySymNames = []string{
		"'+'",
		"'-'",
		"'}'",
		"'\\n'",
		"','",
		"$end",
		"';'",
		"IDENT",
		"'('",
		"'='",
		"FALSE",
		"FUNC",
		"NIL",
		"NUMBER",
		"STRING",
		"TRUE",
		"')'",
		"ELSE",
		"'%'",
		"'*'",
		"'/'",
		"'<'",
		"'>'",
		"EQEQ",
		"GE",
		"LE",
		"NEQ",
		"'{'",
		"expr",
		"IF",
		"RETURN",
		"newLine",
		"newLines",
		"opt_term",
		"term",
		"exprs",
		"program",
		"stmts",
		"opt_newLines",
		"stmt",
		"stmt_if",
		"ident_args",
		"$default",
		"error",
		"UNARY",
	}

	yyTokenLiteralStrings = map[int]string{}

	yyReductions = map[int]struct{ xsym, components int }{
		0:  {0, 1},
		1:  {36, 1},
		2:  {36, 2},
		3:  {37, 2},
		4:  {37, 3},
		5:  {39, 1},
		6:  {39, 3},
		7:  {39, 3},
		8:  {39, 1},
		9:  {40, 5},
		10: {40, 7},
		11: {40, 5},
		12: {40, 2},
		13: {35, 0},
		14: {35, 1},
		15: {35, 4},
		16: {28, 1},
		17: {28, 1},
		18: {28, 1},
		19: {28, 1},
		20: {28, 1},
		21: {28, 1},
		22: {28, 4},
		23: {28, 8},
		24: {28, 2},
		25: {28, 2},
		26: {28, 3},
		27: {28, 3},
		28: {28, 3},
		29: {28, 3},
		30: {28, 3},
		31: {28, 3},
		32: {28, 3},
		33: {28, 3},
		34: {28, 3},
		35: {28, 3},
		36: {28, 3},
		37: {28, 3},
		38: {41, 0},
		39: {41, 1},
		40: {41, 4},
		41: {33, 0},
		42: {33, 1},
		43: {34, 2},
		44: {34, 1},
		45: {34, 1},
		46: {38, 0},
		47: {38, 1},
		48: {32, 1},
		49: {32, 2},
		50: {31, 1},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [93][]uint16{
		// 0
		{10, 10, 3: 59, 10, 10, 56, 10, 10, 10, 10, 10, 10, 10, 10, 10, 29: 10, 10, 58, 57, 53, 55, 36: 52, 54},
		{5: 51},
		{77, 78, 50, 4: 38, 50, 7: 70, 79, 38, 74, 76, 75, 71, 72, 73, 28: 66, 68, 69, 35: 67, 39: 143, 65},
		{2: 10, 59, 5: 10, 56, 31: 58, 57, 62, 63},
		{9, 9, 9, 4: 9, 9, 7: 9, 9, 9, 9, 9, 9, 9, 9, 9, 29: 9, 9},
		// 5
		{6, 6, 6, 59, 6, 6, 7: 6, 6, 6, 6, 6, 6, 6, 6, 6, 29: 6, 6, 58, 61},
		{7, 7, 7, 59, 7, 7, 7: 7, 7, 7, 7, 7, 7, 7, 7, 7, 29: 7, 7, 60},
		{3, 3, 3, 3, 3, 3, 7: 3, 3, 3, 3, 3, 3, 3, 3, 3, 29: 3, 3},
		{1, 1, 1, 1, 1, 1, 7: 1, 1, 1, 1, 1, 1, 1, 1, 1, 29: 1, 1},
		{2, 2, 2, 2, 2, 2, 7: 2, 2, 2, 2, 2, 2, 2, 2, 2, 29: 2, 2},
		// 10
		{8, 8, 8, 59, 8, 8, 7: 8, 8, 8, 8, 8, 8, 8, 8, 8, 29: 8, 8, 60},
		{2: 49, 5: 49},
		{77, 78, 9, 4: 38, 9, 7: 70, 79, 38, 74, 76, 75, 71, 72, 73, 28: 66, 68, 69, 35: 67, 39: 64, 65},
		{2: 47, 47, 5: 47, 47},
		{2: 46, 46, 5: 46, 46, 17: 134},
		// 15
		{82, 83, 43, 43, 37, 43, 43, 9: 132, 18: 86, 84, 85, 90, 89, 87, 91, 92, 88},
		{4: 121, 9: 130},
		{77, 78, 7: 70, 79, 10: 74, 76, 75, 71, 72, 73, 28: 126},
		{77, 78, 38, 38, 38, 38, 38, 70, 79, 10: 74, 76, 75, 71, 72, 73, 17: 38, 28: 119, 35: 125},
		{35, 35, 35, 35, 35, 35, 35, 8: 118, 35, 16: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35},
		// 20
		{34, 34, 34, 34, 34, 34, 34, 9: 34, 16: 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34},
		{33, 33, 33, 33, 33, 33, 33, 9: 33, 16: 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33},
		{32, 32, 32, 32, 32, 32, 32, 9: 32, 16: 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32},
		{31, 31, 31, 31, 31, 31, 31, 9: 31, 16: 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31},
		{30, 30, 30, 30, 30, 30, 30, 9: 30, 16: 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30},
		// 25
		{7: 106},
		{77, 78, 7: 70, 79, 10: 74, 76, 75, 71, 72, 73, 28: 105},
		{77, 78, 7: 70, 79, 10: 74, 76, 75, 71, 72, 73, 28: 104},
		{77, 78, 7: 70, 79, 10: 74, 76, 75, 71, 72, 73, 28: 80},
		{82, 83, 16: 81, 18: 86, 84, 85, 90, 89, 87, 91, 92, 88},
		// 30
		{25, 25, 25, 25, 25, 25, 25, 9: 25, 16: 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25},
		{77, 78, 7: 70, 79, 10: 74, 76, 75, 71, 72, 73, 28: 103},
		{77, 78, 7: 70, 79, 10: 74, 76, 75, 71, 72, 73, 28: 102},
		{77, 78, 7: 70, 79, 10: 74, 76, 75, 71, 72, 73, 28: 101},
		{77, 78, 7: 70, 79, 10: 74, 76, 75, 71, 72, 73, 28: 100},
		// 35
		{77, 78, 7: 70, 79, 10: 74, 76, 75, 71, 72, 73, 28: 99},
		{77, 78, 7: 70, 79, 10: 74, 76, 75, 71, 72, 73, 28: 98},
		{77, 78, 7: 70, 79, 10: 74, 76, 75, 71, 72, 73, 28: 97},
		{77, 78, 7: 70, 79, 10: 74, 76, 75, 71, 72, 73, 28: 96},
		{77, 78, 7: 70, 79, 10: 74, 76, 75, 71, 72, 73, 28: 95},
		// 40
		{77, 78, 7: 70, 79, 10: 74, 76, 75, 71, 72, 73, 28: 94},
		{77, 78, 7: 70, 79, 10: 74, 76, 75, 71, 72, 73, 28: 93},
		{82, 83, 14, 14, 14, 14, 14, 9: 14, 16: 14, 14, 86, 84, 85, 14, 14, 14, 14, 14, 14, 14},
		{82, 83, 15, 15, 15, 15, 15, 9: 15, 16: 15, 15, 86, 84, 85, 15, 15, 15, 15, 15, 15, 15},
		{82, 83, 16, 16, 16, 16, 16, 9: 16, 16: 16, 16, 86, 84, 85, 16, 16, 16, 16, 16, 16, 16},
		// 45
		{82, 83, 17, 17, 17, 17, 17, 9: 17, 16: 17, 17, 86, 84, 85, 17, 17, 17, 17, 17, 17, 17},
		{82, 83, 18, 18, 18, 18, 18, 9: 18, 16: 18, 18, 86, 84, 85, 90, 89, 18, 91, 92, 18, 18},
		{82, 83, 19, 19, 19, 19, 19, 9: 19, 16: 19, 19, 86, 84, 85, 90, 89, 19, 91, 92, 19, 19},
		{20, 20, 20, 20, 20, 20, 20, 9: 20, 16: 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		{21, 21, 21, 21, 21, 21, 21, 9: 21, 16: 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		// 50
		{22, 22, 22, 22, 22, 22, 22, 9: 22, 16: 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		{23, 23, 23, 23, 23, 23, 23, 9: 23, 16: 23, 23, 86, 84, 85, 23, 23, 23, 23, 23, 23, 23},
		{24, 24, 24, 24, 24, 24, 24, 9: 24, 16: 24, 24, 86, 84, 85, 24, 24, 24, 24, 24, 24, 24},
		{26, 26, 26, 26, 26, 26, 26, 9: 26, 16: 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26},
		{27, 27, 27, 27, 27, 27, 27, 9: 27, 16: 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27},
		// 55
		{8: 107},
		{4: 13, 7: 109, 16: 13, 41: 108},
		{4: 111, 16: 110},
		{4: 12, 16: 12},
		{27: 115},
		// 60
		{3: 59, 7: 5, 31: 58, 113, 38: 112},
		{7: 114},
		{4, 4, 3: 59, 7: 4, 4, 10: 4, 4, 4, 4, 4, 4, 31: 60},
		{4: 11, 16: 11},
		{10, 10, 10, 59, 10, 6: 56, 10, 10, 10, 10, 10, 10, 10, 10, 10, 29: 10, 10, 58, 57, 53, 55, 36: 116, 54},
		// 65
		{2: 117},
		{28, 28, 28, 28, 28, 28, 28, 9: 28, 16: 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28},
		{77, 78, 4: 38, 7: 70, 79, 10: 74, 76, 75, 71, 72, 73, 38, 28: 119, 35: 120},
		{82, 83, 37, 37, 37, 37, 37, 16: 37, 37, 86, 84, 85, 90, 89, 87, 91, 92, 88},
		{4: 121, 16: 122},
		// 70
		{5, 5, 3: 59, 7: 5, 5, 10: 5, 5, 5, 5, 5, 5, 31: 58, 113, 38: 123},
		{29, 29, 29, 29, 29, 29, 29, 9: 29, 16: 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29},
		{77, 78, 7: 70, 79, 10: 74, 76, 75, 71, 72, 73, 28: 124},
		{82, 83, 36, 36, 36, 36, 36, 9: 36, 16: 36, 36, 86, 84, 85, 90, 89, 87, 91, 92, 88},
		{2: 39, 39, 121, 39, 39, 17: 39},
		// 75
		{82, 83, 18: 86, 84, 85, 90, 89, 87, 91, 92, 88, 127},
		{10, 10, 10, 59, 10, 6: 56, 10, 10, 10, 10, 10, 10, 10, 10, 10, 29: 10, 10, 58, 57, 53, 55, 36: 128, 54},
		{2: 129},
		{2: 42, 42, 5: 42, 42, 17: 42},
		{77, 78, 38, 38, 38, 38, 38, 70, 79, 10: 74, 76, 75, 71, 72, 73, 28: 119, 35: 131},
		// 80
		{2: 44, 44, 121, 44, 44},
		{77, 78, 7: 70, 79, 10: 74, 76, 75, 71, 72, 73, 28: 133},
		{82, 83, 45, 45, 5: 45, 45, 18: 86, 84, 85, 90, 89, 87, 91, 92, 88},
		{27: 136, 29: 135},
		{77, 78, 7: 70, 79, 10: 74, 76, 75, 71, 72, 73, 28: 139},
		// 85
		{10, 10, 10, 59, 10, 6: 56, 10, 10, 10, 10, 10, 10, 10, 10, 10, 29: 10, 10, 58, 57, 53, 55, 36: 137, 54},
		{2: 138},
		{2: 40, 40, 5: 40, 40, 17: 40},
		{82, 83, 18: 86, 84, 85, 90, 89, 87, 91, 92, 88, 140},
		{10, 10, 10, 59, 10, 6: 56, 10, 10, 10, 10, 10, 10, 10, 10, 10, 29: 10, 10, 58, 57, 53, 55, 36: 141, 54},
		// 90
		{2: 142},
		{2: 41, 41, 5: 41, 41, 17: 41},
		{2: 48, 48, 5: 48, 48},
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
	const yyError = 43

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
			yyVAL.stmt = yyS[yypt-0].stmt_if
		}
	case 6:
		{
			yyVAL.stmt = &ast.AssStmt{Left: []ast.Expr{yyS[yypt-2].expr}, Right: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 7:
		{
			yyVAL.stmt = &ast.AssStmt{Left: yyS[yypt-2].exprs, Right: yyS[yypt-0].exprs}
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
			yyVAL.expr = &ast.UnaryExpr{Operator: "+", Expr: yyS[yypt-0].expr}
		}
	case 25:
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyS[yypt-0].expr}
		}
	case 26:
		{
			yyVAL.expr = &ast.ParentExpr{SubExpr: yyS[yypt-1].expr}
		}
	case 27:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "+", Right: yyS[yypt-0].expr}
		}
	case 28:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "-", Right: yyS[yypt-0].expr}
		}
	case 29:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "*", Right: yyS[yypt-0].expr}
		}
	case 30:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "/", Right: yyS[yypt-0].expr}
		}
	case 31:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "%", Right: yyS[yypt-0].expr}
		}
	case 32:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "==", Right: yyS[yypt-0].expr}
		}
	case 33:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "!=", Right: yyS[yypt-0].expr}
		}
	case 34:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: ">", Right: yyS[yypt-0].expr}
		}
	case 35:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "<", Right: yyS[yypt-0].expr}
		}
	case 36:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: ">=", Right: yyS[yypt-0].expr}
		}
	case 37:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "<=", Right: yyS[yypt-0].expr}
		}
	case 38:
		{
			yyVAL.ident_args = []string{}
		}
	case 39:
		{
			yyVAL.ident_args = []string{yyS[yypt-0].token.Literal}
		}
	case 40:
		{
			yyVAL.ident_args = append(yyS[yypt-3].ident_args, yyS[yypt-0].token.Literal)
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
