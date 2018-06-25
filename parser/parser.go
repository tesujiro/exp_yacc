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
	yyDefault = 57352
	yyEofCode = 57344
	EQEQ      = 57348
	GE        = 57350
	IDENT     = 57346
	LE        = 57351
	NEQ       = 57349
	NUMBER    = 57347
	yyErrCode = 57345

	yyMaxDepth = 200
	yyTabOfs   = -25
)

var (
	yyPrec = map[int]int{
		';':   0,
		'\n':  0,
		'=':   1,
		IDENT: 2,
		EQEQ:  3,
		NEQ:   3,
		'>':   4,
		'<':   4,
		GE:    4,
		LE:    4,
		'+':   5,
		'-':   5,
	}

	yyXLAT = map[int]int{
		57344: 0,  // $end (25x)
		10:    1,  // '\n' (23x)
		40:    2,  // '(' (20x)
		57346: 3,  // IDENT (20x)
		57347: 4,  // NUMBER (20x)
		59:    5,  // ';' (17x)
		43:    6,  // '+' (14x)
		45:    7,  // '-' (14x)
		60:    8,  // '<' (14x)
		62:    9,  // '>' (14x)
		57348: 10, // EQEQ (14x)
		57350: 11, // GE (14x)
		57351: 12, // LE (14x)
		57349: 13, // NEQ (14x)
		41:    14, // ')' (12x)
		61:    15, // '=' (12x)
		57353: 16, // expr (12x)
		57354: 17, // newLine (5x)
		57355: 18, // newLines (3x)
		57356: 19, // opt_term (2x)
		57358: 20, // stmt (2x)
		57360: 21, // term (2x)
		57357: 22, // program (1x)
		57359: 23, // stmts (1x)
		57352: 24, // $default (0x)
		57345: 25, // error (0x)
	}

	yySymNames = []string{
		"$end",
		"'\\n'",
		"'('",
		"IDENT",
		"NUMBER",
		"';'",
		"'+'",
		"'-'",
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
	}

	yyTokenLiteralStrings = map[int]string{}

	yyReductions = map[int]struct{ xsym, components int }{
		0:  {0, 1},
		1:  {22, 2},
		2:  {23, 2},
		3:  {23, 3},
		4:  {20, 3},
		5:  {20, 1},
		6:  {16, 1},
		7:  {16, 1},
		8:  {16, 3},
		9:  {16, 3},
		10: {16, 3},
		11: {16, 3},
		12: {16, 3},
		13: {16, 3},
		14: {16, 3},
		15: {16, 3},
		16: {16, 3},
		17: {19, 0},
		18: {19, 1},
		19: {21, 2},
		20: {21, 1},
		21: {21, 1},
		22: {18, 1},
		23: {18, 2},
		24: {17, 1},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [39][]uint8{
		// 0
		{1: 33, 8, 8, 8, 30, 17: 32, 31, 28, 21: 29, 26, 27},
		{25},
		{8, 33, 5: 30, 17: 32, 31, 61, 21: 62},
		{2: 40, 38, 39, 16: 37, 20: 36},
		{2: 7, 7, 7},
		// 5
		{4, 33, 4, 4, 4, 17: 32, 35},
		{5, 33, 5, 5, 5, 17: 34},
		{3, 3, 3, 3, 3},
		{1, 1, 1, 1, 1},
		{2, 2, 2, 2, 2},
		// 10
		{6, 33, 6, 6, 6, 17: 34},
		{23, 23, 5: 23},
		{20, 20, 5: 20, 43, 44, 48, 47, 45, 49, 50, 46, 15: 59},
		{19, 19, 5: 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19},
		{18, 18, 5: 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18},
		// 15
		{2: 40, 38, 39, 16: 41},
		{6: 43, 44, 48, 47, 45, 49, 50, 46, 42},
		{17, 17, 5: 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
		{2: 40, 38, 39, 16: 58},
		{2: 40, 38, 39, 16: 57},
		// 20
		{2: 40, 38, 39, 16: 56},
		{2: 40, 38, 39, 16: 55},
		{2: 40, 38, 39, 16: 54},
		{2: 40, 38, 39, 16: 53},
		{2: 40, 38, 39, 16: 52},
		// 25
		{2: 40, 38, 39, 16: 51},
		{9, 9, 5: 9, 43, 44, 9, 9, 9, 9, 9, 9, 9, 9},
		{10, 10, 5: 10, 43, 44, 10, 10, 10, 10, 10, 10, 10, 10},
		{11, 11, 5: 11, 43, 44, 11, 11, 11, 11, 11, 11, 11, 11},
		{12, 12, 5: 12, 43, 44, 12, 12, 12, 12, 12, 12, 12, 12},
		// 30
		{13, 13, 5: 13, 43, 44, 48, 47, 13, 49, 50, 13, 13, 13},
		{14, 14, 5: 14, 43, 44, 48, 47, 14, 49, 50, 14, 14, 14},
		{15, 15, 5: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		{16, 16, 5: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		{2: 40, 38, 39, 16: 60},
		// 35
		{21, 21, 5: 21, 43, 44, 48, 47, 45, 49, 50, 46},
		{24},
		{7, 2: 40, 38, 39, 16: 37, 20: 63},
		{22, 22, 5: 22},
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
	const yyError = 25

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
			yyVAL.expr = ast.ParentExpr{SubExpr: yyS[yypt-1].expr}
		}
	case 9:
		{
			yyVAL.expr = ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "+", Right: yyS[yypt-0].expr}
		}
	case 10:
		{
			yyVAL.expr = ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "-", Right: yyS[yypt-0].expr}
		}
	case 11:
		{
			yyVAL.expr = ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "==", Right: yyS[yypt-0].expr}
		}
	case 12:
		{
			yyVAL.expr = ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "!=", Right: yyS[yypt-0].expr}
		}
	case 13:
		{
			yyVAL.expr = ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: ">", Right: yyS[yypt-0].expr}
		}
	case 14:
		{
			yyVAL.expr = ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "<", Right: yyS[yypt-0].expr}
		}
	case 15:
		{
			yyVAL.expr = ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: ">=", Right: yyS[yypt-0].expr}
		}
	case 16:
		{
			yyVAL.expr = ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "<=", Right: yyS[yypt-0].expr}
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
