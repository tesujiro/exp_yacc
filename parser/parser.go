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
	yyTabOfs   = -50
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
		43:    0,  // '+' (65x)
		45:    1,  // '-' (65x)
		125:   2,  // '}' (53x)
		10:    3,  // '\n' (49x)
		57344: 4,  // $end (47x)
		59:    5,  // ';' (40x)
		57346: 6,  // IDENT (40x)
		40:    7,  // '(' (38x)
		57350: 8,  // FALSE (36x)
		57352: 9,  // FUNC (36x)
		57351: 10, // NIL (36x)
		57347: 11, // NUMBER (36x)
		57348: 12, // STRING (36x)
		57349: 13, // TRUE (36x)
		44:    14, // ',' (32x)
		41:    15, // ')' (31x)
		57359: 16, // ELSE (30x)
		37:    17, // '%' (29x)
		42:    18, // '*' (29x)
		47:    19, // '/' (29x)
		60:    20, // '<' (29x)
		62:    21, // '>' (29x)
		57354: 22, // EQEQ (29x)
		57356: 23, // GE (29x)
		57357: 24, // LE (29x)
		57355: 25, // NEQ (29x)
		123:   26, // '{' (26x)
		61:    27, // '=' (23x)
		57362: 28, // expr (22x)
		57358: 29, // IF (15x)
		57353: 30, // RETURN (14x)
		57365: 31, // newLine (12x)
		57366: 32, // newLines (9x)
		57368: 33, // opt_term (6x)
		57373: 34, // term (6x)
		57369: 35, // program (5x)
		57372: 36, // stmts (5x)
		57363: 37, // exprs (2x)
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
		"$end",
		"';'",
		"IDENT",
		"'('",
		"FALSE",
		"FUNC",
		"NIL",
		"NUMBER",
		"STRING",
		"TRUE",
		"','",
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
		"'='",
		"expr",
		"IF",
		"RETURN",
		"newLine",
		"newLines",
		"opt_term",
		"term",
		"program",
		"stmts",
		"exprs",
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
		1:  {35, 1},
		2:  {35, 2},
		3:  {36, 2},
		4:  {36, 3},
		5:  {39, 1},
		6:  {39, 3},
		7:  {39, 1},
		8:  {40, 5},
		9:  {40, 7},
		10: {40, 5},
		11: {40, 2},
		12: {37, 0},
		13: {37, 1},
		14: {37, 4},
		15: {28, 1},
		16: {28, 1},
		17: {28, 1},
		18: {28, 1},
		19: {28, 1},
		20: {28, 1},
		21: {28, 4},
		22: {28, 8},
		23: {28, 2},
		24: {28, 2},
		25: {28, 3},
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
		37: {41, 0},
		38: {41, 1},
		39: {41, 4},
		40: {33, 0},
		41: {33, 1},
		42: {34, 2},
		43: {34, 1},
		44: {34, 1},
		45: {38, 0},
		46: {38, 1},
		47: {32, 1},
		48: {32, 2},
		49: {31, 1},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [90][]uint16{
		// 0
		{10, 10, 3: 58, 10, 55, 10, 10, 10, 10, 10, 10, 10, 10, 29: 10, 10, 57, 56, 52, 54, 51, 53},
		{4: 50},
		{75, 76, 49, 4: 49, 6: 68, 77, 72, 74, 73, 69, 70, 71, 28: 65, 66, 67, 39: 139, 64},
		{2: 10, 58, 10, 55, 31: 57, 56, 61, 62},
		{9, 9, 9, 4: 9, 6: 9, 9, 9, 9, 9, 9, 9, 9, 29: 9, 9},
		// 5
		{6, 6, 6, 58, 6, 6: 6, 6, 6, 6, 6, 6, 6, 6, 29: 6, 6, 57, 60},
		{7, 7, 7, 58, 7, 6: 7, 7, 7, 7, 7, 7, 7, 7, 29: 7, 7, 59},
		{3, 3, 3, 3, 3, 6: 3, 3, 3, 3, 3, 3, 3, 3, 29: 3, 3},
		{1, 1, 1, 1, 1, 6: 1, 1, 1, 1, 1, 1, 1, 1, 29: 1, 1},
		{2, 2, 2, 2, 2, 6: 2, 2, 2, 2, 2, 2, 2, 2, 29: 2, 2},
		// 10
		{8, 8, 8, 58, 8, 6: 8, 8, 8, 8, 8, 8, 8, 8, 29: 8, 8, 59},
		{2: 48, 4: 48},
		{75, 76, 9, 4: 9, 6: 68, 77, 72, 74, 73, 69, 70, 71, 28: 65, 66, 67, 39: 63, 64},
		{2: 46, 46, 46, 46},
		{2: 45, 45, 45, 45, 16: 130},
		// 15
		{80, 81, 43, 43, 43, 43, 17: 84, 82, 83, 88, 87, 85, 89, 90, 86, 27: 128},
		{75, 76, 6: 68, 77, 72, 74, 73, 69, 70, 71, 28: 124},
		{75, 76, 38, 38, 38, 38, 68, 77, 72, 74, 73, 69, 70, 71, 38, 16: 38, 28: 117, 37: 123},
		{35, 35, 35, 35, 35, 35, 7: 116, 14: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35},
		{34, 34, 34, 34, 34, 34, 14: 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34},
		// 20
		{33, 33, 33, 33, 33, 33, 14: 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33},
		{32, 32, 32, 32, 32, 32, 14: 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32},
		{31, 31, 31, 31, 31, 31, 14: 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31},
		{30, 30, 30, 30, 30, 30, 14: 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30},
		{6: 104},
		// 25
		{75, 76, 6: 68, 77, 72, 74, 73, 69, 70, 71, 28: 103},
		{75, 76, 6: 68, 77, 72, 74, 73, 69, 70, 71, 28: 102},
		{75, 76, 6: 68, 77, 72, 74, 73, 69, 70, 71, 28: 78},
		{80, 81, 15: 79, 17: 84, 82, 83, 88, 87, 85, 89, 90, 86},
		{25, 25, 25, 25, 25, 25, 14: 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25},
		// 30
		{75, 76, 6: 68, 77, 72, 74, 73, 69, 70, 71, 28: 101},
		{75, 76, 6: 68, 77, 72, 74, 73, 69, 70, 71, 28: 100},
		{75, 76, 6: 68, 77, 72, 74, 73, 69, 70, 71, 28: 99},
		{75, 76, 6: 68, 77, 72, 74, 73, 69, 70, 71, 28: 98},
		{75, 76, 6: 68, 77, 72, 74, 73, 69, 70, 71, 28: 97},
		// 35
		{75, 76, 6: 68, 77, 72, 74, 73, 69, 70, 71, 28: 96},
		{75, 76, 6: 68, 77, 72, 74, 73, 69, 70, 71, 28: 95},
		{75, 76, 6: 68, 77, 72, 74, 73, 69, 70, 71, 28: 94},
		{75, 76, 6: 68, 77, 72, 74, 73, 69, 70, 71, 28: 93},
		{75, 76, 6: 68, 77, 72, 74, 73, 69, 70, 71, 28: 92},
		// 40
		{75, 76, 6: 68, 77, 72, 74, 73, 69, 70, 71, 28: 91},
		{80, 81, 14, 14, 14, 14, 14: 14, 14, 14, 84, 82, 83, 14, 14, 14, 14, 14, 14, 14, 14},
		{80, 81, 15, 15, 15, 15, 14: 15, 15, 15, 84, 82, 83, 15, 15, 15, 15, 15, 15, 15, 15},
		{80, 81, 16, 16, 16, 16, 14: 16, 16, 16, 84, 82, 83, 16, 16, 16, 16, 16, 16, 16, 16},
		{80, 81, 17, 17, 17, 17, 14: 17, 17, 17, 84, 82, 83, 17, 17, 17, 17, 17, 17, 17, 17},
		// 45
		{80, 81, 18, 18, 18, 18, 14: 18, 18, 18, 84, 82, 83, 88, 87, 18, 89, 90, 18, 18, 18},
		{80, 81, 19, 19, 19, 19, 14: 19, 19, 19, 84, 82, 83, 88, 87, 19, 89, 90, 19, 19, 19},
		{20, 20, 20, 20, 20, 20, 14: 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		{21, 21, 21, 21, 21, 21, 14: 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		{22, 22, 22, 22, 22, 22, 14: 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		// 50
		{23, 23, 23, 23, 23, 23, 14: 23, 23, 23, 84, 82, 83, 23, 23, 23, 23, 23, 23, 23, 23},
		{24, 24, 24, 24, 24, 24, 14: 24, 24, 24, 84, 82, 83, 24, 24, 24, 24, 24, 24, 24, 24},
		{26, 26, 26, 26, 26, 26, 14: 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26},
		{27, 27, 27, 27, 27, 27, 14: 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27},
		{7: 105},
		// 55
		{6: 107, 14: 13, 13, 41: 106},
		{14: 109, 108},
		{14: 12, 12},
		{26: 113},
		{3: 58, 6: 5, 31: 57, 111, 38: 110},
		// 60
		{6: 112},
		{4, 4, 3: 58, 6: 4, 4, 4, 4, 4, 4, 4, 4, 31: 59},
		{14: 11, 11},
		{10, 10, 10, 58, 5: 55, 10, 10, 10, 10, 10, 10, 10, 10, 29: 10, 10, 57, 56, 52, 54, 114, 53},
		{2: 115},
		// 65
		{28, 28, 28, 28, 28, 28, 14: 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28},
		{75, 76, 6: 68, 77, 72, 74, 73, 69, 70, 71, 38, 38, 28: 117, 37: 118},
		{80, 81, 37, 37, 37, 37, 14: 37, 37, 37, 84, 82, 83, 88, 87, 85, 89, 90, 86},
		{14: 119, 120},
		{5, 5, 3: 58, 6: 5, 5, 5, 5, 5, 5, 5, 5, 31: 57, 111, 38: 121},
		// 70
		{29, 29, 29, 29, 29, 29, 14: 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29},
		{75, 76, 6: 68, 77, 72, 74, 73, 69, 70, 71, 28: 122},
		{80, 81, 36, 36, 36, 36, 14: 36, 36, 36, 84, 82, 83, 88, 87, 85, 89, 90, 86},
		{2: 39, 39, 39, 39, 14: 119, 16: 39},
		{80, 81, 17: 84, 82, 83, 88, 87, 85, 89, 90, 86, 125},
		// 75
		{10, 10, 10, 58, 5: 55, 10, 10, 10, 10, 10, 10, 10, 10, 29: 10, 10, 57, 56, 52, 54, 126, 53},
		{2: 127},
		{2: 42, 42, 42, 42, 16: 42},
		{75, 76, 6: 68, 77, 72, 74, 73, 69, 70, 71, 28: 129},
		{80, 81, 44, 44, 44, 44, 17: 84, 82, 83, 88, 87, 85, 89, 90, 86},
		// 80
		{26: 132, 29: 131},
		{75, 76, 6: 68, 77, 72, 74, 73, 69, 70, 71, 28: 135},
		{10, 10, 10, 58, 5: 55, 10, 10, 10, 10, 10, 10, 10, 10, 29: 10, 10, 57, 56, 52, 54, 133, 53},
		{2: 134},
		{2: 40, 40, 40, 40, 16: 40},
		// 85
		{80, 81, 17: 84, 82, 83, 88, 87, 85, 89, 90, 86, 136},
		{10, 10, 10, 58, 5: 55, 10, 10, 10, 10, 10, 10, 10, 10, 29: 10, 10, 57, 56, 52, 54, 137, 53},
		{2: 138},
		{2: 41, 41, 41, 41, 16: 41},
		{2: 47, 47, 47, 47},
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
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr}
		}
	case 8:
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-3].expr, Then: yyS[yypt-1].stmts, Else: nil}
		}
	case 9:
		{
			yyVAL.stmt_if.(*ast.IfStmt).ElseIf = append(yyVAL.stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyS[yypt-3].expr, Then: yyS[yypt-1].stmts})
		}
	case 10:
		{
			if yyVAL.stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				//$$.(*ast.IfStmt).Else = append($$.(*ast.IfStmt).Else, $4...)
				yyVAL.stmt_if.(*ast.IfStmt).Else = yyS[yypt-1].stmts
			}
		}
	case 11:
		{
			yyVAL.stmt_if = &ast.ReturnStmt{Exprs: yyS[yypt-0].exprs}
		}
	case 12:
		{
			yyVAL.exprs = nil
		}
	case 13:
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 14:
		{
			yyVAL.exprs = append(yyS[yypt-3].exprs, yyS[yypt-0].expr)
		}
	case 15:
		{
			yyVAL.expr = &ast.IdentExpr{Literal: yyS[yypt-0].token.Literal}
		}
	case 16:
		{
			yyVAL.expr = &ast.NumExpr{Literal: yyS[yypt-0].token.Literal}
		}
	case 17:
		{
			yyVAL.expr = &ast.StringExpr{Literal: yyS[yypt-0].token.Literal}
		}
	case 18:
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyS[yypt-0].token.Literal}
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
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].token.Literal, SubExprs: yyS[yypt-1].exprs}
		}
	case 22:
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-6].token.Literal, Args: yyS[yypt-4].ident_args, Stmts: yyS[yypt-1].stmts}
		}
	case 23:
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "+", Expr: yyS[yypt-0].expr}
		}
	case 24:
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyS[yypt-0].expr}
		}
	case 25:
		{
			yyVAL.expr = &ast.ParentExpr{SubExpr: yyS[yypt-1].expr}
		}
	case 26:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "+", Right: yyS[yypt-0].expr}
		}
	case 27:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "-", Right: yyS[yypt-0].expr}
		}
	case 28:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "*", Right: yyS[yypt-0].expr}
		}
	case 29:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "/", Right: yyS[yypt-0].expr}
		}
	case 30:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "%", Right: yyS[yypt-0].expr}
		}
	case 31:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "==", Right: yyS[yypt-0].expr}
		}
	case 32:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "!=", Right: yyS[yypt-0].expr}
		}
	case 33:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: ">", Right: yyS[yypt-0].expr}
		}
	case 34:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "<", Right: yyS[yypt-0].expr}
		}
	case 35:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: ">=", Right: yyS[yypt-0].expr}
		}
	case 36:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "<=", Right: yyS[yypt-0].expr}
		}
	case 37:
		{
			yyVAL.ident_args = []string{}
		}
	case 38:
		{
			yyVAL.ident_args = []string{yyS[yypt-0].token.Literal}
		}
	case 39:
		{
			yyVAL.ident_args = append(yyS[yypt-3].ident_args, yyS[yypt-0].token.Literal)
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
