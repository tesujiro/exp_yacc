// Code generated by goyacc - DO NOT EDIT.

package parser

import __yyfmt__ "fmt"

import "github.com/tesujiro/exp_yacc/ast"

type yySymType struct {
	yys   int
	token ast.Token
	expr  ast.Expr
	stmt  ast.Stmt
	stmts []ast.Stmt
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault = 57353
	yyEofCode = 57344
	EQEQ      = 57348
	GE        = 57350
	IDENT     = 57346
	LE        = 57351
	NEQ       = 57349
	NUMBER    = 57347
	UNARY     = 57352
	yyErrCode = 57345

	yyMaxDepth = 200
	yyTabOfs   = -30
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
		43:    0,  // '+' (44x)
		45:    1,  // '-' (44x)
		57344: 2,  // $end (30x)
		10:    3,  // '\n' (28x)
		40:    4,  // '(' (25x)
		57346: 5,  // IDENT (25x)
		57347: 6,  // NUMBER (25x)
		59:    7,  // ';' (22x)
		37:    8,  // '%' (19x)
		42:    9,  // '*' (19x)
		47:    10, // '/' (19x)
		60:    11, // '<' (19x)
		62:    12, // '>' (19x)
		57348: 13, // EQEQ (19x)
		57350: 14, // GE (19x)
		57351: 15, // LE (19x)
		57349: 16, // NEQ (19x)
		41:    17, // ')' (17x)
		61:    18, // '=' (17x)
		57354: 19, // expr (17x)
		57355: 20, // newLine (5x)
		57356: 21, // newLines (3x)
		57357: 22, // opt_term (2x)
		57359: 23, // stmt (2x)
		57361: 24, // term (2x)
		57358: 25, // program (1x)
		57360: 26, // stmts (1x)
		57353: 27, // $default (0x)
		57345: 28, // error (0x)
		57352: 29, // UNARY (0x)
	}

	yySymNames = []string{
		"'+'",
		"'-'",
		"$end",
		"'\\n'",
		"'('",
		"IDENT",
		"NUMBER",
		"';'",
		"'%'",
		"'*'",
		"'/'",
		"'<'",
		"'>'",
		"EQEQ",
		"GE",
		"LE",
		"NEQ",
		"')'",
		"'='",
		"expr",
		"newLine",
		"newLines",
		"opt_term",
		"stmt",
		"term",
		"program",
		"stmts",
		"$default",
		"error",
		"UNARY",
	}

	yyTokenLiteralStrings = map[int]string{}

	yyReductions = map[int]struct{ xsym, components int }{
		0:  {0, 1},
		1:  {25, 2},
		2:  {26, 2},
		3:  {26, 3},
		4:  {23, 3},
		5:  {23, 1},
		6:  {19, 1},
		7:  {19, 1},
		8:  {19, 2},
		9:  {19, 2},
		10: {19, 3},
		11: {19, 3},
		12: {19, 3},
		13: {19, 3},
		14: {19, 3},
		15: {19, 3},
		16: {19, 3},
		17: {19, 3},
		18: {19, 3},
		19: {19, 3},
		20: {19, 3},
		21: {19, 3},
		22: {22, 0},
		23: {22, 1},
		24: {24, 2},
		25: {24, 1},
		26: {24, 1},
		27: {21, 1},
		28: {21, 2},
		29: {20, 1},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [49][]uint8{
		// 0
		{8, 8, 3: 38, 8, 8, 8, 35, 20: 37, 36, 33, 24: 34, 31, 32},
		{2: 30},
		{2: 8, 38, 7: 35, 20: 37, 36, 76, 24: 77},
		{45, 46, 4: 47, 43, 44, 19: 42, 23: 41},
		{7, 7, 4: 7, 7, 7},
		// 5
		{4, 4, 4, 38, 4, 4, 4, 20: 37, 40},
		{5, 5, 5, 38, 5, 5, 5, 20: 39},
		{3, 3, 3, 3, 3, 3, 3},
		{1, 1, 1, 1, 1, 1, 1},
		{2, 2, 2, 2, 2, 2, 2},
		// 10
		{6, 6, 6, 38, 6, 6, 6, 20: 39},
		{2: 28, 28, 7: 28},
		{50, 51, 25, 25, 7: 25, 54, 52, 53, 58, 57, 55, 59, 60, 56, 18: 74},
		{24, 24, 24, 24, 7: 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24},
		{23, 23, 23, 23, 7: 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23},
		// 15
		{45, 46, 4: 47, 43, 44, 19: 73},
		{45, 46, 4: 47, 43, 44, 19: 72},
		{45, 46, 4: 47, 43, 44, 19: 48},
		{50, 51, 8: 54, 52, 53, 58, 57, 55, 59, 60, 56, 49},
		{20, 20, 20, 20, 7: 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		// 20
		{45, 46, 4: 47, 43, 44, 19: 71},
		{45, 46, 4: 47, 43, 44, 19: 70},
		{45, 46, 4: 47, 43, 44, 19: 69},
		{45, 46, 4: 47, 43, 44, 19: 68},
		{45, 46, 4: 47, 43, 44, 19: 67},
		// 25
		{45, 46, 4: 47, 43, 44, 19: 66},
		{45, 46, 4: 47, 43, 44, 19: 65},
		{45, 46, 4: 47, 43, 44, 19: 64},
		{45, 46, 4: 47, 43, 44, 19: 63},
		{45, 46, 4: 47, 43, 44, 19: 62},
		// 30
		{45, 46, 4: 47, 43, 44, 19: 61},
		{50, 51, 9, 9, 7: 9, 54, 52, 53, 9, 9, 9, 9, 9, 9, 9, 9},
		{50, 51, 10, 10, 7: 10, 54, 52, 53, 10, 10, 10, 10, 10, 10, 10, 10},
		{50, 51, 11, 11, 7: 11, 54, 52, 53, 11, 11, 11, 11, 11, 11, 11, 11},
		{50, 51, 12, 12, 7: 12, 54, 52, 53, 12, 12, 12, 12, 12, 12, 12, 12},
		// 35
		{50, 51, 13, 13, 7: 13, 54, 52, 53, 58, 57, 13, 59, 60, 13, 13, 13},
		{50, 51, 14, 14, 7: 14, 54, 52, 53, 58, 57, 14, 59, 60, 14, 14, 14},
		{15, 15, 15, 15, 7: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		{16, 16, 16, 16, 7: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		{17, 17, 17, 17, 7: 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
		// 40
		{18, 18, 18, 18, 7: 18, 54, 52, 53, 18, 18, 18, 18, 18, 18, 18, 18},
		{19, 19, 19, 19, 7: 19, 54, 52, 53, 19, 19, 19, 19, 19, 19, 19, 19},
		{21, 21, 21, 21, 7: 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		{22, 22, 22, 22, 7: 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		{45, 46, 4: 47, 43, 44, 19: 75},
		// 45
		{50, 51, 26, 26, 7: 26, 54, 52, 53, 58, 57, 55, 59, 60, 56},
		{2: 29},
		{45, 46, 7, 4: 47, 43, 44, 19: 42, 23: 78},
		{2: 27, 27, 7: 27},
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
	const yyError = 28

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
			yyVAL.stmts = yyS[yypt-1].stmts
			yylex.(*Lexer).Result = yyVAL.stmts
		}
	case 2:
		{
			yyVAL.stmts = []ast.Stmt{yyS[yypt-0].stmt}
		}
	case 3:
		{
			yyVAL.stmts = append(yyS[yypt-2].stmts, yyS[yypt-0].stmt)
		}
	case 4:
		{
			yyVAL.stmt = ast.AssStmt{Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 5:
		{
			yyVAL.stmt = ast.ExprStmt{Expr: yyS[yypt-0].expr}
		}
	case 6:
		{
			yyVAL.expr = ast.IdentExpr{Literal: yyS[yypt-0].token.Literal}
		}
	case 7:
		{
			yyVAL.expr = ast.NumExpr{Literal: yyS[yypt-0].token.Literal}
		}
	case 8:
		{
			yyVAL.expr = ast.UnaryExpr{Operator: "+", Expr: yyS[yypt-0].expr}
		}
	case 9:
		{
			yyVAL.expr = ast.UnaryExpr{Operator: "-", Expr: yyS[yypt-0].expr}
		}
	case 10:
		{
			yyVAL.expr = ast.ParentExpr{SubExpr: yyS[yypt-1].expr}
		}
	case 11:
		{
			yyVAL.expr = ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "+", Right: yyS[yypt-0].expr}
		}
	case 12:
		{
			yyVAL.expr = ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "-", Right: yyS[yypt-0].expr}
		}
	case 13:
		{
			yyVAL.expr = ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "*", Right: yyS[yypt-0].expr}
		}
	case 14:
		{
			yyVAL.expr = ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "/", Right: yyS[yypt-0].expr}
		}
	case 15:
		{
			yyVAL.expr = ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "%", Right: yyS[yypt-0].expr}
		}
	case 16:
		{
			yyVAL.expr = ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "==", Right: yyS[yypt-0].expr}
		}
	case 17:
		{
			yyVAL.expr = ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "!=", Right: yyS[yypt-0].expr}
		}
	case 18:
		{
			yyVAL.expr = ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: ">", Right: yyS[yypt-0].expr}
		}
	case 19:
		{
			yyVAL.expr = ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "<", Right: yyS[yypt-0].expr}
		}
	case 20:
		{
			yyVAL.expr = ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: ">=", Right: yyS[yypt-0].expr}
		}
	case 21:
		{
			yyVAL.expr = ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "<=", Right: yyS[yypt-0].expr}
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
